use crate::proto::vectordb::Database;

/// Defines the core database management operations.
/// Implement this trait for different database backends.
pub trait DatabaseManager: Send + Sync + 'static {
    /// Creates a new instance of the database manager.
    fn new() -> Self where Self: Sized;
    
    /// Lists all available databases.
    fn list_databases(&self) -> Vec<Database>;
    
    /// Retrieves a specific database by name.
    fn get_database(&self, name: &str) -> Option<&Database>;
    
    /// Adds a new database.
    fn add_database(&mut self, database: Database) -> Result<(), String>;
}
