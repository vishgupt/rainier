//! Collection module containing core collection abstractions and implementations.

// Traits
pub mod collection_manager;

// Implementations
pub mod inmemory_collection_manager;

// Re-export for cleaner imports
pub use collection_manager::CollectionManager;
pub use inmemory_collection_manager::InMemoryCollectionManager;
