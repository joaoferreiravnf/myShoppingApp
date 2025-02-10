import React, { useEffect, useState } from "react";

const ItemList = () => {
    const [items, setItems] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        // Fetch items from the backend
        fetch("http://localhost:8080/items") // Adjust the endpoint if needed
            .then((response) => {
                if (!response.ok) {
                    throw new Error("Network response was not ok");
                }
                return response.json();
            })
            .then((data) => {
                setItems(data);
                setLoading(false);
            })
            .catch((error) => {
                setError(error);
                setLoading(false);
            });
    }, []);

    if (loading) return <p>Loading...</p>;
    if (error) return <p>Error: {error.message}</p>;

    return (
        <table>
            <thead>
            <tr>
                <th>Name</th>
                <th>Qty</th>
                <th>Type</th>
                <th>Market</th>
                <th>When</th>
                <th>Who</th>
                <th>Actions</th>
            </tr>
            </thead>
            <tbody>
            {items.map((item, index) => (
                <tr key={index}>
                    <td>{item.name}</td>
                    <td>{item.quantity}</td>
                    <td>{item.type}</td>
                    <td>{item.market}</td>
                    <td>{item.when}</td>
                    <td>{item.who}</td>
                    <td>
                        <button
                            onClick={() => {
                                handleDelete(item.id);
                            }}
                            style={{ backgroundColor: "red", color: "white" }}
                        >
                            Delete
                        </button>
                    </td>
                </tr>
            ))}
            </tbody>
        </table>
    );
};

const handleDelete = (id) => {
    fetch(`/items/${id}`, { method: "DELETE" })
        .then((response) => {
            if (!response.ok) {
                throw new Error("Failed to delete item");
            }
            alert("Item deleted successfully!");
            window.location.reload(); // Reload the page
        })
        .catch((error) => {
            console.error("Error deleting item:", error);
        });
};

export default ItemList;