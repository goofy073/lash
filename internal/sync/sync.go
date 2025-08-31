package sync

import (
	"context"
	"log/slog"
	"remnawave-tg-shop-bot/internal/channel"
	"remnawave-tg-shop-bot/internal/database"
	"time"
)

type SyncService struct {
	channelManager     *channel.Manager
	customerRepository *database.CustomerRepository
}

func NewSyncService(channelManager *channel.Manager, customerRepository *database.CustomerRepository) *SyncService {
	return &SyncService{
		channelManager: channelManager, customerRepository: customerRepository,
	}
}

// Sync checks for expired subscriptions and removes users from channel
func (s SyncService) Sync() {
	slog.Info("Starting channel access sync")
	ctx := context.Background()
	
	// Find all customers with expired subscriptions
	now := time.Now()
	expiredCustomers, err := s.customerRepository.FindByExpirationRange(ctx, time.Time{}, now)
	if err != nil {
		slog.Error("Error while getting expired customers", "error", err)
		return
	}

	if expiredCustomers == nil || len(*expiredCustomers) == 0 {
		slog.Info("No expired customers found")
		return
	}

	var processedCount int
	for _, customer := range *expiredCustomers {
		// Skip if customer never had channel access
		if customer.ChannelInviteLink == nil || !customer.IsInChannel {
			continue
		}

		// Ban user from channel
		err := s.channelManager.BanUser(ctx, customer.TelegramID)
		if err != nil {
			slog.Error("Failed to ban expired user", "telegramId", customer.TelegramID, "error", err)
			continue
		}

		// Update customer status
		updates := map[string]interface{}{
			"is_in_channel": false,
		}
		err = s.customerRepository.UpdateFields(ctx, customer.ID, updates)
		if err != nil {
			slog.Error("Failed to update customer status", "customerId", customer.ID, "error", err)
			continue
		}

		// Revoke the invite link if it still exists
		if customer.ChannelInviteLink != nil && *customer.ChannelInviteLink != "" {
			err = s.channelManager.RevokeInviteLink(ctx, *customer.ChannelInviteLink)
			if err != nil {
				slog.Error("Failed to revoke invite link", "link", *customer.ChannelInviteLink, "error", err)
			}
		}

		processedCount++
	}

	slog.Info("Channel access sync completed", "processedExpired", processedCount)
}
