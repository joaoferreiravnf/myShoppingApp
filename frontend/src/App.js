import React, { useState, useEffect } from "react";
import ItemTable from "./components/ItemTable";
import AddItemForm from "./components/AddItemForm";

const App = () => {
    const { isAuthenticated, isLoading } = useAuth0();
    if (isLoading) return <p>Loading...</p>;
    if (!isAuthenticated) {
        return <LoginPage />;
    }

    return (
        <div className="container">
            <h1>Items</h1>
            <ItemTable />
            <AddItemForm />
        </div>
    );
    const [items, setItems] = useState([]);
    const [markets, setMarkets] = useState([]);
    const [types, setTypes] = useState([]);
    const [quantities, setQuantities] = useState([]);
    const LogoutButton = () => {
        const { logout } = useAuth0();

        return (
            <button onClick={() => logout({ logoutParams: { returnTo: window.location.origin } })}>
                Log Out
            </button>
        );
    };

    useEffect(() => {
        const fetchItems = async () => {
            try {
                const response = await fetch("http://localhost:8080/items");
                const data = await response.json();

                // Populate state from JSON response
                setItems(data.items);
                setMarkets(data.markets);
                setTypes(data.types);
                setQuantities(data.quantities);
            } catch (error) {
                console.error("Error fetching items:", error);
            }
        };

        fetchItems();
    }, []);

    const handleAddItem = async (newItem) => {
        // Add new item (example POST request)
        const response = await fetch("http://localhost:8080/items", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(newItem),
        });

        const addedItem = await response.json();
        setItems([...items, addedItem]);
    };

    return (
        <div className="container">
            <h1>Items</h1>
            <ItemTable items={items} />
            <AddItemForm onAdd={handleAddItem} markets={markets} types={types} quantities={quantities} />
        </div>
    );
};

export default App;