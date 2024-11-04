// src/ItemsList.js
import React, { useState, useEffect } from 'react';

function ItemsList() {
    const [items, setItems] = useState([]);       // State to store items from the backend
    const [loading, setLoading] = useState(true);  // Loading state for user feedback
    const [error, setError] = useState(null);      // Error state to handle issues

    useEffect(() => {
        // Function to fetch items from the backend
        async function fetchItems() {
            try {
                const response = await fetch("http://localhost:8080/items");
                if (!response.ok) {
                    throw new Error("Failed to fetch items");
                }

                const data = await response.json();
                setItems(data); // Set the fetched items to state
            } catch (err) {
                setError(err.message);
            } finally {
                setLoading(false); // Loading complete
            }
        }

        fetchItems();
    }, []); // Empty dependency array means this runs only on component mount

    // Display a loading message, error message, or the list of items
    if (loading) return <p>Loading items...</p>;
    if (error) return <p>Error: {error}</p>;

    return (
        <div>
            <h1>Items List</h1>
            <ul>
                {items.map(item => (
                    <li key={item.id}>
                        <strong>{item.Name}</strong> - {item.Quantity} - {item.Market} - {item.AddedAt}
                    </li>
                ))}
            </ul>
        </div>
    );
}

export default ItemsList;