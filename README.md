# Stuart Go Client

A Go client library for accessing the Stuart API, a same-day delivery service platform.

## Overview

This client provides access to [Stuart's API](https://api-docs.stuart.com/), enabling integration with their delivery network for pickup and dropoff services across multiple cities.

## Main Components

### Core Client
- **ClientWrapper**: Main client with OAuth2 authentication
- **ClientInterface**: Defines all available API operations
- **NewClient()**: Creates authenticated client for sandbox or production environments

### Key Features

**Job Management**
- Create, update, cancel, and retrieve delivery jobs
- Get job pricing estimates and validate job parameters
- Retrieve ETA for pickup and dropoff
- Schedule jobs with time slots

**Delivery Operations**
- Cancel specific deliveries
- Get driver phone numbers for active deliveries
- Track delivery status and proof of delivery

**Address & Zone Services**
- Validate pickup/dropoff addresses
- Get zone coverage areas (GeoJSON format)
- Find nearby parcel shops with schedules

**Package Types**: XSmall, Small, Medium, Large, XLarge
**Restricted Items**: Tobacco, alcohol, knives, fireworks, etc.

## API Functionality

The Stuart API supports:
- **Real-time delivery management** with live tracking
- **Flexible scheduling** with time windows
- **Multiple transport types** (bike, car, van)
- **Zone-based coverage** across supported cities
- **Proof of delivery** with signatures and photos
- **Pricing transparency** with tax calculations
- **Webhook notifications** for status updates

## Usage

```go
client := NewClient(ctx, stuartclient.SandboxEnv, apiClientId, apiClientSecret)
job, err := client.CreateJob(ctx, jobRequest)
```

Supports both sandbox and production environments for testing and deployment.
