package channel

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"remnawave-tg-shop-bot/internal/config"
	"remnawave-tg-shop-bot/utils"
)

type Manager struct {
	bot *bot.Bot
}

func NewManager(bot *bot.Bot) *Manager {
	return &Manager{
		bot: bot,
	}
}

// GenerateInviteLink creates a new invite link for the private channel
func (m *Manager) GenerateInviteLink(ctx context.Context, telegramID int64, expireDate time.Time) (string, error) {
	// Create invite link that expires at the subscription end date
	inviteLink, err := m.bot.CreateChatInviteLink(ctx, &bot.CreateChatInviteLinkParams{
		ChatID:     config.PrivateChannelId(),
		Name:       fmt.Sprintf("User_%d", telegramID),
		ExpireDate: &expireDate,
		MemberLimit: 1, // Only allow one use of this invite link
	})
	
	if err != nil {
		slog.Error("Failed to create invite link", "telegramId", utils.MaskHalfInt64(telegramID), "error", err)
		return "", fmt.Errorf("failed to create invite link: %w", err)
	}

	slog.Info("Created channel invite link", "telegramId", utils.MaskHalfInt64(telegramID), "expireDate", expireDate)
	return inviteLink.InviteLink, nil
}

// RevokeInviteLink revokes an existing invite link
func (m *Manager) RevokeInviteLink(ctx context.Context, inviteLink string) error {
	_, err := m.bot.RevokeChatInviteLink(ctx, &bot.RevokeChatInviteLinkParams{
		ChatID:     config.PrivateChannelId(),
		InviteLink: inviteLink,
	})
	
	if err != nil {
		slog.Error("Failed to revoke invite link", "inviteLink", inviteLink, "error", err)
		return fmt.Errorf("failed to revoke invite link: %w", err)
	}

	slog.Info("Revoked channel invite link", "inviteLink", inviteLink)
	return nil
}

// BanUser removes a user from the channel (when subscription expires)
func (m *Manager) BanUser(ctx context.Context, telegramID int64) error {
	_, err := m.bot.BanChatMember(ctx, &bot.BanChatMemberParams{
		ChatID:    config.PrivateChannelId(),
		UserID:    telegramID,
		UntilDate: 0, // Permanent ban (they can be unbanned if they subscribe again)
	})
	
	if err != nil {
		slog.Error("Failed to ban user from channel", "telegramId", utils.MaskHalfInt64(telegramID), "error", err)
		return fmt.Errorf("failed to ban user from channel: %w", err)
	}

	slog.Info("Banned user from channel", "telegramId", utils.MaskHalfInt64(telegramID))
	return nil
}

// UnbanUser unbans a user from the channel (when they renew subscription)
func (m *Manager) UnbanUser(ctx context.Context, telegramID int64) error {
	_, err := m.bot.UnbanChatMember(ctx, &bot.UnbanChatMemberParams{
		ChatID:        config.PrivateChannelId(),
		UserID:        telegramID,
		OnlyIfBanned: true,
	})
	
	if err != nil {
		slog.Error("Failed to unban user from channel", "telegramId", utils.MaskHalfInt64(telegramID), "error", err)
		return fmt.Errorf("failed to unban user from channel: %w", err)
	}

	slog.Info("Unbanned user from channel", "telegramId", utils.MaskHalfInt64(telegramID))
	return nil
}

// CheckChannelMembership verifies if a user is a member of the channel
func (m *Manager) CheckChannelMembership(ctx context.Context, telegramID int64) (bool, error) {
	member, err := m.bot.GetChatMember(ctx, &bot.GetChatMemberParams{
		ChatID: config.PrivateChannelId(),
		UserID: telegramID,
	})
	
	if err != nil {
		slog.Error("Failed to check channel membership", "telegramId", utils.MaskHalfInt64(telegramID), "error", err)
		return false, fmt.Errorf("failed to check channel membership: %w", err)
	}

	// Check if user is an active member
	switch member.Status {
	case models.ChatMemberStatusCreator, models.ChatMemberStatusAdministrator, models.ChatMemberStatusMember:
		return true, nil
	case models.ChatMemberStatusRestricted:
		// Check if restricted member can still access the channel
		if restricted, ok := member.MemberInfo.(models.ChatMemberRestricted); ok {
			return !restricted.IsKicked, nil
		}
		return false, nil
	default:
		return false, nil
	}
}

// CreateTrialInviteLink creates a trial invite link with limited duration
func (m *Manager) CreateTrialInviteLink(ctx context.Context, telegramID int64, trialDays int) (string, error) {
	expireDate := time.Now().UTC().AddDate(0, 0, trialDays)
	
	inviteLink, err := m.bot.CreateChatInviteLink(ctx, &bot.CreateChatInviteLinkParams{
		ChatID:     config.PrivateChannelId(),
		Name:       fmt.Sprintf("Trial_User_%d", telegramID),
		ExpireDate: &expireDate,
		MemberLimit: 1,
	})
	
	if err != nil {
		slog.Error("Failed to create trial invite link", "telegramId", utils.MaskHalfInt64(telegramID), "error", err)
		return "", fmt.Errorf("failed to create trial invite link: %w", err)
	}

	slog.Info("Created trial channel invite link", "telegramId", utils.MaskHalfInt64(telegramID), "trialDays", trialDays)
	return inviteLink.InviteLink, nil
}