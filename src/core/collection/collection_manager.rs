use crate::proto::vectordb::Collection;

/// Defines the core collection management operations.
/// Implement this trait for different collection backends.
pub trait CollectionManager: Send + Sync + 'static {
    /// Creates a new instance of the collection manager.
    fn new() -> Self where Self: Sized;
    
    /// Lists all collections in a database.
    fn list_collections(&self, database_name: &str) -> Vec<Collection>;
    
    /// Retrieves a specific collection by name.
    fn get_collection(&self, database_name: &str, collection_name: &str) -> Option<&Collection>;
    
    /// Adds a new collection to a database.
    fn add_collection(&mut self, collection: Collection) -> Result<(), String>;
}
