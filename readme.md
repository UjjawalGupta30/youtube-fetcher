# YouTube API - Fampay Hiring

## Tech Stack
- **Golang**: An open-source programming language designed for simplicity and performance.
- **MongoDB**: A NoSQL database used for storing fetched video data.
- **Gin**: A fast and lightweight web framework for creating APIs.
- **YouTube Data API v3**: Used for fetching the latest video data directly from YouTube.

## Features
- **Continuous Video Fetching**: Automatically fetches videos from YouTube API and stores them in a MongoDB database.
- **Paginated Video Retrieval**: Provides an API to retrieve stored videos in reverse chronological order of their publishing date with pagination.
- **API Key Rotation**: Implements API key rotation to manage YouTube API quota limits efficiently.

## Setup

1. **Clone the repository**
   ```bash
   git clone https://github.com/yourusername/youtube-fetcher.git
   cd youtube-fetcher
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Create an `.env` file and add the following variables**
   ```plaintext
   YOUTUBE_API_KEYS="<key1>,<key2>,..."
   MONGO_URI="mongodb+srv://<username>:<password>@<cluster>/<database>?retryWrites=true&w=majority"
   QUERY_INTERVAL=10
   ```

4. **Run the server**
   ```bash
   go run main.go
   ```

## API Endpoints

### Get Videos
Retrieve a paginated list of videos stored in the database, sorted by publishing date in reverse chronological order:
```bash
GET /videos
```
- **Query Parameters**:
  - `page`: The page number (default: 1).
  - `pageSize`: Number of videos per page (default: 10).

Example:
```bash
GET /videos?page=2&pageSize=4
![image](https://github.com/user-attachments/assets/e67bd7b5-0131-4635-a695-afcc9ca2fe03)

```
