# Telegram Channel Access Bot
[![Stars](https://img.shields.io/github/stars/Jolymmiels/remnawave-telegram-shop.svg?style=social)](https://github.com/Jolymmiels/remnawave-telegram-shop/stargazers)
[![Forks](https://img.shields.io/github/forks/Jolymmiels/remnawave-telegram-shop.svg?style=social)](https://github.com/Jolymmiels/remnawave-telegram-shop/network/members)
[![Issues](https://img.shields.io/github/issues/Jolymmiels/remnawave-telegram-shop.svg)](https://github.com/Jolymmiels/remnawave-telegram-shop/issues)
## Description

A Telegram bot for selling access to private Telegram channels. This service allows users to purchase and manage channel access subscriptions through Telegram with multiple payment system options.

## Admin commands

- `/sync` - Clean up expired channel memberships and remove users who no longer have active subscriptions.

## Features

- Purchase private channel access with different payment methods (bank cards, cryptocurrency, Telegram Stars)
- Multiple subscription plans (1, 3, 6, 12 months)
- Automated channel access management with invite links
- **Subscription Notifications**: The bot automatically sends notifications to users 3 days before their channel access expires
- Multi-language support (Russian and English)
- **Trial Access**: Allow users to try the channel for free for a limited time
- **Referral System**: Users can invite friends and get bonus access time
- **Automatic User Management**: Expired users are automatically removed from the channel
- All telegram message support HTML formatting https://core.telegram.org/bots/api#html-style
## Environment Variables

The application requires the following environment variables to be set:

| Variable                 | Description                                                                                                                                  |
|--------------------------|----------------------------------------------------------------------------------------------------------------------------------------------| 
| `TELEGRAM_TOKEN`         | Telegram Bot API token for bot functionality                                                                                                 |
| `ADMIN_TELEGRAM_ID`      | Admin telegram id                                                                                                                            |
| `PRIVATE_CHANNEL_ID`     | The ID or username of the private channel (e.g., @my_private_channel or -1001234567890)                                                      |
| `PRIVATE_CHANNEL_USERNAME` | The username of the private channel (without @)                                                                                              |
| `DATABASE_URL`           | PostgreSQL connection string                                                                                                                 |
| `PRICE_1`                | Price for 1 month access (in kopecks)                                                                                                        |
| `PRICE_3`                | Price for 3 months access (in kopecks)                                                                                                       |
| `PRICE_6`                | Price for 6 months access (in kopecks)                                                                                                       |
| `PRICE_12`               | Price for 12 months access (in kopecks)                                                                                                      |
| `STARS_PRICE_1`          | Price in Telegram Stars for 1 month (optional)                                                                                               |
| `STARS_PRICE_3`          | Price in Telegram Stars for 3 months (optional)                                                                                              |
| `STARS_PRICE_6`          | Price in Telegram Stars for 6 months (optional)                                                                                              |
| `STARS_PRICE_12`         | Price in Telegram Stars for 12 months (optional)                                                                                             |
| `TRIAL_DAYS`             | Number of days for trial access (set to 0 to disable)                                                                                        |
| `REFERRAL_DAYS`          | Number of bonus days for referrals (set to 0 to disable)                                                                                     |
| `CRYPTO_PAY_ENABLED`     | Enable/disable CryptoPay payment method (true/false)                                                                                         |
| `CRYPTO_PAY_TOKEN`       | CryptoPay API token (required if CryptoPay enabled)                                                                                          |
| `CRYPTO_PAY_URL`         | CryptoPay API URL (required if CryptoPay enabled)                                                                                            |
| `YOOKASA_ENABLED`        | Enable/disable YooKassa payment method (true/false)                                                                                          |
| `YOOKASA_SECRET_KEY`     | YooKassa API secret key (required if YooKassa enabled)                                                                                       |
| `YOOKASA_SHOP_ID`        | YooKassa shop identifier (required if YooKassa enabled)                                                                                      |
| `YOOKASA_URL`            | YooKassa API URL (required if YooKassa enabled)                                                                                              |
| `YOOKASA_EMAIL`          | Email address associated with YooKassa account (required if YooKassa enabled)                                                                |
| `TELEGRAM_STARS_ENABLED` | Enable/disable Telegram Stars payment method (true/false)                                                                                    |
| `SERVER_STATUS_URL`      | URL to server status page (optional) - if not set, button will not be displayed                                                              |
| `SUPPORT_URL`            | URL to support chat or page (optional) - if not set, button will not be displayed                                                            |
| `FEEDBACK_URL`           | URL to feedback/reviews page (optional) - if not set, button will not be displayed                                                           |
| `CHANNEL_URL`            | URL to public Telegram channel (optional) - if not set, button will not be displayed                                                         |
| `TOS_URL`                | URL to Terms of Service page (optional) - if not set, button will not be displayed                                                           |
| `MINI_APP_URL`           | Telegram Web App URL (optional) - if not set, regular inline buttons will be used                                                            |

## User Interface

The bot dynamically creates buttons based on available environment variables:

- Main buttons for purchasing and accessing the private channel are always shown
- Additional buttons for Server Status, Support, Feedback, and Public Channel are only displayed if their corresponding URL environment variables are set

## Automated Notifications

The bot includes a notification system that runs daily at 16:00 UTC to check for expiring channel access:

- Users receive a notification 3 days before their channel access expires
- The notification includes the exact expiration date and a convenient button to renew access
- Notifications are sent in the user's preferred language

## Channel Management

The bot automatically manages channel access:

- Creates unique invite links for each customer with expiration dates
- Automatically removes users from the channel when their subscription expires
- Supports trial access with limited time periods
- Handles referral bonuses by extending access time

## Bot Setup Requirements

**Important**: Your bot must be an **administrator** in the private channel with the following permissions:
- Can invite users
- Can restrict members
- Can delete messages (optional, for cleanup)

Without these permissions, the bot cannot manage channel access properly.

## Plugins and Dependencies

### Telegram Bot

- [Telegram Bot API](https://core.telegram.org/bots/api)
- [Go Telegram Bot API](https://github.com/go-telegram/bot)

### Database

- [PostgreSQL](https://www.postgresql.org/)
- [pgx - PostgreSQL Driver](https://github.com/jackc/pgx)

### Payment Systems

- [YooKassa API](https://yookassa.ru/developers/api)
- [CryptoPay API](https://help.crypt.bot/crypto-pay-api)
- Telegram Stars

## Setup Instructions

1. Clone the repository

```bash
git clone https://github.com/Jolymmiels/remnawave-telegram-shop && cd remnawave-telegram-shop
```

2. Create a `.env` file in the root directory with all the environment variables listed above

```bash
mv .env.sample .env
```

3. Run the bot:

   ```bash
   docker compose up -d
   ```

## How to change bot messages

Go to folder translations inside bot folder and change needed language.

## Update Instructions

1. Pull the latest Docker image:

   ```bash
   docker compose pull
   ```


2. Restart the containers:
   ```bash
   docker compose down && docker compose up -d
   ```

## Donations

If you appreciate this project and want to help keep it running (and fuel those caffeine-fueled coding marathons),
consider donating. Your support helps drive future updates and improvements.

**Donation Methods:**

- **Bep20 USDT:** `0x4D1ee2445fdC88fA49B9d02FB8ee3633f45Bef48`

- **SOL Solana:** `HNQhe6SCoU5UDZicFKMbYjQNv9Muh39WaEWbZayQ9Nn8`

- **TRC20 USDT:** `TBJrguLia8tvydsQ2CotUDTYtCiLDA4nPW`

- **TON USDT:** `UQAdAhVxOr9LS07DDQh0vNzX2575Eu0eOByjImY1yheatXgr`

