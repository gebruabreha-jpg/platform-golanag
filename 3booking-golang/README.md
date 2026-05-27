run:
go clean -cache -testcache -modcache
go clean
go mod tidy
gofmt -w .
golangci-lint run
go build ./cmd/myapp
go run main.go
go test ./...

Language: Go 1.21 (backend)
Tech Stack:
Backend: Go 1.21 with Gin web framework, GORM ORM
Frontend: Next.js 14 + TypeScript (seen in frontend/ directory)
Database: PostgreSQL
Cache: Redis
Authentication: JWT
API Documentation: Swagger/OpenAPI
Deployment: Docker + Docker Compose
Testing: Go testing framework, Jest for frontend
Note: Despite the directory name "4connectme11java" suggesting Java, the actual implementation uses Go for the backend. The project connects Ethiopians globally through 4 core services: Flight Booking, Peer-to-Peer Shipping, Rental Marketplace, and Service Marketplace (for professionals like lawyers, doctors, etc.)

Deployments:-
Frontend:
Vercel (best for Next.js - zero config, automatic deployments)
Netlify
Cloudflare Pages


Backend:
Fly.io (recommended) - Multi-region, good for Go apps, handles multiple services
Render.com - Docker Compose support, managed PostgreSQL
Railway.app - Simple setup

Serverless Database:
Supabase - PostgreSQL-compatible with auth, realtime, storage
PlanetScale - MySQL-compatible (not suitable for this project)
Neon.tech - Serverless PostgreSQL

Best Free Database Options:-
Supabase - Free tier: 500MB database, 1GB storage, includes auth/realtime
Neon.tech - Free tier: 10GB storage, 1 compute unit, 30 million rows
Railway.app - Free tier: 500MB PostgreSQL (simple but limited)
Render.com - Free tier: 100MB PostgreSQL (very limited)
AWS RDS - Free tier: 20GB PostgreSQL for 12 months (then paid)


all-in-one: Render.com (Docker Compose deployment with managed PostgreSQL)


## Core Services:
1. **Flight Booking Service** - Book international flights mainly to ethiopia.
2. **Peer-to-Peer Shipping & Travel Delivery** - Send/receive items via travelers
3. **Rental Marketplace** - Rooms/accomidation with habesh community
4. **Service Marketplace** - Lawyers, doctors, and skilled professionals

---
## 🌍 Our Services
### 1. Flight Booking Service
Book Ethiopian airline tickets and international flights through trusted agents within the community.
**Features:**
- Search and compare flight prices from multiple airlines
- Secure payment processing for ticket purchases
- E-ticket delivery and itinerary management
- Flight change and cancellation assistance
- Travel insurance options
- Multi-city and group booking capabilities
 
**Supported Services:**
- Ethiopian Airlines ticket booking
- International flight reservations
- Domestic flights within Africa
- Charter and private flight arrangements
- Travel package combinations (flight + hotel)


### 2. Peer-to-Peer Shipping & Travel Delivery (Parcel/Last-Mile Delivery Platform)
 
Connect with travelers heading to Ethiopia to send/receive items through unused luggage space. This operates as a parcel/delivery marketplace and last-mile delivery platform.
**Features:**
- Find travelers matching your route & dates automatically
- Secure peer-to-peer (p2p) escrow payments (payment released on delivery confirmation)
- Real-time package tracking
- Verified traveler profiles with trust scores
- Community ratings & reviews

**Inspired by:** PiggyBee, Grabr, Worldee, Airlift Hub
---

### 3: Housing (Medium Priority)
**Use Case 3: Find Habesha Roommate in USA and EUROP**
- User searches DC housing → filters for Habesha preferences → contacts landlord → books viewing → signs lease
- Estimated: 2 weeks

### 4. Service Marketplace(lowst priority)
Connect with verified professionals including lawyers, doctors, consultants, and skilled service providers.
**Features:**
- Verified professional profiles with specialization tags
- Service categories: legal, medical, consulting, technical, creative, and more
- Consultation booking system
- Client reviews and ratings
- Remote service options
 
**Types of Professionals:**
- Legal: lawyers, immigration consultants, paralegals
- Medical: doctors, specialists, therapists, counselors
- Business: consultants, accountants, financial advisors
- Technical: IT specialists, engineers, developers
- Creative: designers, writers, translators, artists
- Skilled trades: contractors, electricians, plumbers, mechanics
---
