# Nilo Commerce
**Closed-loop B2B2C Cross-Border Bill Payment Platform**
> **Legal Disclaimer**: This is a *bill payment platform*, NOT a remittance service. We operate under the "agent of the payee" exception per FinCEN ruling FIN-2008-R006, allowing us to legally process payments on behalf of merchants without requiring money transmitter licenses.

## 📋 Table of Contents
- [Business Model](#business-model)
- [How It Works](#how-it-works)
- [Architecture](#architecture)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [API Documentation](#api-documentation)
- [Deployment](#deployment)
- [Compliance & Legal](#compliance--legal)
- [Contributing](#contributing)
- [License](#license)

## 💼 Business Model
Nilo Commerce connects the **Ethiopian diaspora** with local merchants (schools, hospitals, supermarkets) through a secure closed-loop bill payment system.

### Value Proposition
- **For Diaspora**: Peace of mind knowing funds go directly to essentials, no cash diversion risk
- **For Merchants**: Eliminate unpaid bills and credit risk with guaranteed upfront payments
- **For Platform**: Transaction convenience fee + FX spread revenue

### Revenue Model
1. **Convenience Fee**: Charged to sender (diaspora) via Stripe (typically 2-3%)
2. **FX Spread**: Difference between mid-market rate and payout rate (typically 50-100 bps)

### Compliance Structure
- Operates as **"Agent of the Payee"** (merchant payment processor)
- NOT a money transmitter - funds flow directly from payer to merchant via closed-loop
- No holding of customer funds - immediate settlement to merchant
- Fully compliant with FinCEN FIN-2008-R006 ruling

## 🔄 How It Works

```
Diaspora User → Stripe (USD) → Nilo Commerce (Platform) → Wise (FX conversion) → Ethiopian Merchant (ETB)
```

### Step-by-Step Flow

**1. Order Creation (Diaspora)**
- User selects merchant from directory (school, hospital, supermarket)
- Enters amount in USD, service type (tuition, healthcare, grocery voucher)
- Completes payment via Stripe (credit/debit card, Apple Pay)

**2. Payment Confirmation**
- Stripe sends webhook to our backend confirming successful payment
- Transaction status updates to `PROCESSING`

**3. Payout via Wise**
- Our backend initiates Wise Business transfer to merchant's Ethiopian bank account
- Wise converts USD → ETB at current rate
- Merchant receives local bank transfer (CBE, Awash, etc.)

**4. Notification**
- Merchant receives SMS/email confirmation of payment
- Diaspora user receives receipt and order confirmation

### Wise Business Integration (Critical)

**Account Structure**:
- ✅ **Do**: Keep USD balance in Wise Business account
- ✅ **Do**: Accumulate funds and cash out quarterly (not weekly)
- ❌ **Don't**: Trigger frequent high-risk transfers to Ethiopia
- ❌ **Don't**: Use personal Wise account for business

**Why This Keeps Your Account Safe**:
- Zero high-risk outbound transfers from Wise's perspective
- Money just sits safely in USD (low compliance risk)
- Less frequent transfers = far fewer fraud triggers

## 🏗️ Architecture

```
┌─────────────┐     ┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│   Flutter   │     │  Next.js    │     │  Backend    │     │   Third-    │
│  Mobile App │     │   Web App   │     │  (Node.js)  │     │   Party     │
│   (iOS/     │     │  (Vercel)   │     │  (AWS/ECS)  │     │  APIs       │
│   Android)  │     │             │     │             │     │             │
└──────┬──────┘     └──────┬──────┘     └──────┬──────┘     └──────┬──────┘
       │                   │                   │                   │
       │ REST API          │                   │                   │
       ├──────────────────►│                   │                   │
       │                   │                   │                   │
       │   Webhooks        │                   │      Webhooks     │
       │◄──────────────────┤                   │◄──────────────────┤
       │                   │                   │                   │
                       ┌───┴───────────────────┴───────────────────┴───┐
                       │          PostgreSQL (Supabase/AWS RDS)         │
                       └─────────────────────────────────────────────────┘
```

### Data Flow
1. Client apps (Flutter/Next.js) → Backend REST API
2. Backend processes transaction → Creates Stripe payment intent
3. Stripe handles payment → Webhook to backend
4. Backend triggers Wise payout
5. Wise transfers to merchant bank account
6. Backend updates transaction status

## 🛠️ Tech Stack

| Layer | Technology | Purpose |
|-------|------------|---------|
| **Mobile** | Flutter | iOS & Android apps (single codebase) |
| **Web** | Next.js + TypeScript | SEO-friendly merchant & user dashboards |
| **Backend** | Node.js + TypeScript | REST API with Express |
| **Database** | PostgreSQL + Prisma | ACID-compliant transactions |
| **Cache** | Redis (optional) | Session storage & caching |
| **Auth** | JWT + bcrypt | Secure authentication |
| **Payments** | Stripe | International card payments |
| **Payouts** | Wise Business API | Cross-border currency conversion & transfers |
| **Hosting** | AWS ECS/Docker | Backend containerization |
| **Web Host** | Vercel | Next.js deployment |

### Key Dependencies (Backend)
- `express` - Web server framework
- `prisma` - ORM for PostgreSQL
- `stripe` - Payment processing
- `@wise-community/wise-api` - Payouts
- `jsonwebtoken` - Authentication
- `winston` - Logging
- `joi` - Input validation

### Key Dependencies (Frontend)
- `flutter_bloc` - State management (Flutter)
- `stripe_sdk` / `flutter_stripe` - Payment UI
- `dio` - HTTP client
- `zustand` - State management (Next.js)
- `@stripe/react-stripe-js` - Stripe Elements

## 📁 Project Structure

```
niloecommerce/
├── backend/                    # Node.js + TypeScript API
│   ├── src/
│   │   ├── controllers/        # Request handlers
│   │   ├── services/           # Business logic & third-party integrations
│   │   ├── models/             # Type definitions (Prisma models)
│   │   ├── middleware/         # Auth, validation, rate limiting
│   │   ├── routes/             # API route definitions
│   │   ├── utils/              # Helpers (logger, validator, etc.)
│   │   ├── config/             # Database & app config
│   │   └── server.ts           # Entry point
│   ├── prisma/
│   │   └── schema.prisma       # Database schema
│   ├── tests/                  # Unit & integration tests
│   ├── docs/                   # API documentation
│   ├── package.json
│   └── tsconfig.json
├── mobile/                     # Flutter app
│   ├── lib/
│   │   ├── core/              # App-wide constants, theme, routing
│   │   ├── data/              # Repositories, data sources
│   │   ├── domain/            # Entities, use cases
│   │   ├── presentation/      # UI screens & widgets
│   │   └── utils/             # Helpers
│   ├── assets/
│   ├── pubspec.yaml
│   └── README.md
├── web/                        # Next.js web app
│   ├── app/                   # App router pages (Next.js 13+)
│   ├── components/            # Reusable UI components
│   ├── lib/                   # Utilities, API client, stores
│   ├── public/                # Static assets
│   ├── styles/                # Global CSS + Tailwind
│   ├── package.json
│   └── tailwind.config.ts
├── docs/                       # Project documentation
├── .github/                    # CI/CD workflows
├── docker-compose.yml          # Local development environment
├── .env.example               # Environment variables template
└── README.md                  # This file
```

## 🚀 Getting Started

### Prerequisites

- Node.js 20+
- Flutter 3.0+ (optional, for mobile)
- Docker & Docker Compose (optional, for local DB)
- PostgreSQL (if not using Docker)
- Stripe account (test mode)
- Wise Business account

### Quick Start (Development)

**1. Clone & Install**

```bash
git clone <your-repo>
cd niloecommerce
```

**2. Setup Environment**

```bash
# Copy example env file
cp .env.example .env

# Edit .env with your API keys
nano .env  # or use any text editor
```

**3. Start Database**

```bash
# Using Docker (recommended)
docker-compose up -d postgres redis

# OR manually install PostgreSQL
# Ensure database 'nilo_commerce' exists
```

**4. Run Backend**

```bash
cd backend
npm install
npx prisma generate
npx prisma db push

# Start dev server
npm run dev
# Server runs at http://localhost:3000
```

**5. Run Frontend (Web)**

```bash
cd web
npm install
npm run dev
# App available at http://localhost:3001
```

**6. Run Flutter App**

```bash
cd mobile
flutter pub get
flutter run
# Runs on connected device/emulator
```

## 🧑‍💻 Development Setup

### Database Migrations

```bash
cd backend

# Create migration
npx prisma migrate dev --name migration_name

# Apply migrations
npx prisma db push

# Generate Prisma client
npx prisma generate

# Open Prisma Studio (GUI)
npx prisma studio
```

### API Testing

Use Postman or curl:

```bash
# Health check
curl http://localhost:3000/health

# Register
curl -X POST http://localhost:3000/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123","firstName":"Test","lastName":"User"}'

# Login
curl -X POST http://localhost:3000/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123"}'
```

### Testing Stripe Webhooks Locally

```bash
# Install Stripe CLI
# https://stripe.com/docs/stripe-cli

# Forward events to localhost
stripe listen --forward-to localhost:3000/api/webhooks/stripe

# In another terminal, trigger a test event
stripe trigger payment_intent.succeeded
```

### Testing Wise Webhooks

Similar to Stripe, use ngrok to expose localhost:

```bash
ngrok http 3000
# Set webhook URL in Wise dashboard to https://your-ngrok-url.ngrok.io/api/webhooks/wise
```

## 📚 API Documentation

### Base URL
```
Production: https://api.nilocommerce.com
Development: http://localhost:3000
```

### Authentication

All protected endpoints require a Bearer token:

```http
Authorization: Bearer <your_jwt_token>
```

### Endpoints

#### Auth
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/auth/register` | Register new user |
| POST | `/api/auth/login` | Login user |
| POST | `/api/auth/logout` | Logout (invalidate token) |
| GET | `/api/auth/me` | Get current user profile |

#### Transactions
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/transactions` | List user's transactions |
| GET | `/api/transactions/:id` | Get transaction details |
| POST | `/api/transactions` | Create new payment |
| POST | `/api/transactions/:id/confirm` | Confirm payment (used by webhook) |

#### Merchants
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/merchants` | List verified merchants |
| GET | `/api/merchants/:id` | Get merchant details |

#### Payouts
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/payouts` | List payouts |
| GET | `/api/payouts/:id/status` | Check payout status |

#### Webhooks
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/webhooks/stripe` | Stripe events |
| POST | `/api/webhooks/wise` | Wise transfer status |

### Response Format

Success:
```json
{
  "success": true,
  "data": { ... }
}
```

Error:
```json
{
  "error": "Error message",
  "details": {} // optional
}
```

## 🚢 Deployment

### Backend (AWS ECS / Docker)

```dockerfile
# backend/Dockerfile
FROM node:20-alpine AS builder
WORKDIR /app
COPY package*.json ./
RUN npm ci
COPY . .
RUN npm run build

FROM node:20-alpine
WORKDIR /app
COPY --from=builder /app/dist ./dist
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/package*.json ./

EXPOSE 3000
CMD ["node", "dist/server.js"]
```

```bash
# Build & push
docker build -t nilocommerce/backend:latest backend/
docker push nilocommerce/backend:latest

# Deploy to ECS (example)
aws ecs update-service --cluster nilo-cluster --service backend --force-new-deployment
```

### Web (Vercel)

```bash
cd web
vercel --prod
```

### Mobile (App Store / Google Play)

```bash
cd mobile
flutter build apk --release   # Android
flutter build ios --release  # iOS (requires Mac)
```

## ⚖️ Compliance & Legal

### "Agent of the Payee" Exception

Under FinCEN ruling FIN-2008-R006, we are exempt from money transmitter regulations because:

- We act as an authorized agent of the merchant
- Payment discharges consumer's debt immediately upon receipt
- We never hold or transmit customer funds
- Funds are immediately used to settle the specific bill

### Prohibited Language

**ALWAYS USE**:
- ✅ "E-commerce bill payment platform"
- ✅ "Direct merchant payment system"
- ✅ "Closed-loop payment solution"

**NEVER USE** (will get account banned):
- ❌ "Remittance"
- ❌ "Sending money home"
- ❌ "Money transfer"
- ❌ "Pool funding"

### Required Documentation

1. **Business Registration**: US LLC (Delaware or Wyoming recommended)
2. **Wise Business Account**: With proper business description
3. **Stripe Business Account**: With accurate business category
4. **Merchant Agreements**: Signed contracts with each school/hospital/supermarket
5. **KYC/AML Program**: Documented compliance procedures

## 🔒 Security Best Practices

### Platform
- All API endpoints require HTTPS
- Rate limiting: 100 requests/15min per IP
- JWT tokens expire in 24 hours
- Passwords hashed with bcrypt (12 rounds)
- Input validation with Joi on all endpoints

### Stripe
- Use only Stripe.js (never handle raw card data)
- Webhook signatures verified
- PCI DSS compliance via Stripe

### Wise
- API token stored in environment variables
- IP whitelisting on Wise dashboard
- Transfer limits enforced

## 📊 Monitoring & Logging

- All log output goes to `logs/` directory
- Winston logger with file + console transports
- Sensitive data never logged
- Audit log tracks all user actions

## 🧪 Testing

```bash
cd backend

# Unit tests
npm test

# Watch mode
npm run test:watch

# Integration tests
npm run test:integration

# Type checking
npm run typecheck

# Linting
npm run lint
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

Please read [CONTRIBUTING.md](docs/CONTRIBUTING.md) for detailed guidelines.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 📞 Support

- **Email**: support@nilocommerce.com
- **Documentation**: https://docs.nilocommerce.com
- **Issues**: https://github.com/your-org/niloecommerce/issues

---

© 2026 Nilo Commerce. Empowering Ethiopian families through secure direct payments.
