# Rainier Vector Database

A high-performance vector database built in Rust with REST APIs, designed for similarity search and vector operations.

## Features

- **REST API**: Complete HTTP API following industry best practices from Pinecone and Qdrant
- **Collection Management**: Create, list, get, and delete vector collections
- **Vector Operations**: Upsert, retrieve, and delete vectors with metadata
- **Similarity Search**: Single and batch vector search with filtering support
- **Type Safety**: Full Rust type system with JSON serialization
- **High Performance**: Built on Axum async web framework
- **CORS Support**: Ready for web client integration

## API Endpoints

### Collection Management
- `POST /collections` - Create a new collection
- `GET /collections` - List all collections
- `GET /collections/{name}` - Get collection details
- `DELETE /collections/{name}` - Delete a collection

### Vector Operations
- `POST /collections/{name}/vectors/upsert` - Add or update vectors
- `GET /collections/{name}/vectors?ids=...` - Get vectors by IDs
- `DELETE /collections/{name}/vectors` - Delete vectors

### Search
- `POST /collections/{name}/search` - Vector similarity search
- `POST /collections/{name}/search/batch` - Batch search queries

### System
- `GET /health` - Health check

## Quick Start

1. **Clone the repository**:
   ```bash
   git clone https://github.com/vishgupt/rainier.git
   cd rainier
   ```

2. **Build and run**:
   ```bash
   cargo run
   ```

3. **Test the API**:
   ```bash
   # Health check
   curl http://localhost:3000/health
   
   # Create a collection
   curl -X POST http://localhost:3000/collections \
     -H "Content-Type: application/json" \
     -d '{"name": "my-collection", "dimension": 768, "metric": "cosine"}'
   
   # List collections
   curl http://localhost:3000/collections
   ```

## Example Usage

### Create a Collection
```bash
curl -X POST http://localhost:3000/collections \
  -H "Content-Type: application/json" \
  -d '{
    "name": "documents",
    "dimension": 1536,
    "metric": "cosine",
    "index_config": {
      "index_type": "hnsw",
      "m": 16,
      "ef_construct": 200
    }
  }'
```

### Upsert Vectors
```bash
curl -X POST http://localhost:3000/collections/documents/vectors/upsert \
  -H "Content-Type: application/json" \
  -d '{
    "vectors": [
      {
        "id": "doc1",
        "values": [0.1, 0.2, 0.3, ...],
        "metadata": {"title": "Document 1", "category": "tech"}
      }
    ]
  }'
```

### Search Vectors
```bash
curl -X POST http://localhost:3000/collections/documents/search \
  -H "Content-Type: application/json" \
  -d '{
    "vector": [0.1, 0.2, 0.3, ...],
    "top_k": 10,
    "filter": {"category": {"$eq": "tech"}},
    "include_metadata": true
  }'
```

## Architecture

- **Framework**: Axum (async web framework)
- **Serialization**: Serde for JSON handling
- **Type System**: Strongly typed with Rust structs
- **Error Handling**: Structured error responses
- **Middleware**: CORS support, extensible for auth/logging

## Development Status

This is the initial API framework implementation. The following components are planned for future development:

- [ ] Vector storage backend (in-memory/persistent)
- [ ] Vector indexing algorithms (HNSW, IVF)
- [ ] Similarity search implementation
- [ ] Metadata filtering engine
- [ ] Persistence and recovery
- [ ] Authentication and authorization
- [ ] Metrics and monitoring

## Contributing

Contributions are welcome! Please feel free to submit issues and pull requests.

## License

This project is open source. Please see the LICENSE file for details.
