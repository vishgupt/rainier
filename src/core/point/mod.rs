//! Point module containing core vector/point abstractions and implementations.

// Traits
pub mod point_manager;

// Implementations
pub mod inmemory_point_manager;

// Re-export for cleaner imports
pub use point_manager::PointManager;
pub use inmemory_point_manager::InMemoryPointManager;
