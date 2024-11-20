const API_URL = "http://localhost:5000/items";

export const fetchItems = async () => {
    const response = await fetch(API_URL);
    return response.json();
};

export const createItem = async (item) => {
    const response = await fetch(API_URL, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(item),
    });
    return response.json();
};

export const deleteItem = async (id) => {
    await fetch(`${API_URL}/${id}`, {
        method: "DELETE",
    });
};