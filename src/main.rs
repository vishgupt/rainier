mod types;
mod handlers;
mod routes;

use routes::create_router;

#[tokio::main]
async fn main() {
    // Initialize the router
    let app = create_router();

    // Define the server address
    let listener = tokio::net::TcpListener::bind("0.0.0.0:3000")
        .await
        .expect("Failed to bind to address");

    println!("🚀 Rainier Vector Database Server starting on http://0.0.0.0:3000");
    println!("📋 Available endpoints:");
    println!("  GET  /health                                    - Health check");
    println!("  POST /collections                               - Create collection");
    println!("  GET  /collections                               - List collections");
    println!("  GET  /collections/{{name}}                        - Get collection info");
    println!("  DELETE /collections/{{name}}                     - Delete collection");
    println!("  POST /collections/{{name}}/vectors/upsert        - Upsert vectors");
    println!("  GET  /collections/{{name}}/vectors?ids=...       - Get vectors");
    println!("  DELETE /collections/{{name}}/vectors             - Delete vectors");
    println!("  POST /collections/{{name}}/search                - Search vectors");
    println!("  POST /collections/{{name}}/search/batch          - Batch search");

    // Start the server
    axum::serve(listener, app)
        .await
        .expect("Failed to start server");
}
