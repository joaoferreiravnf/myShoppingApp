import React from "react";

const ItemTable = ({ items, handleDelete }) => {
    return (
        <table className="table table-striped table-hover">
            <thead>
            <tr>
                <th>Name</th>
                <th>Quantity</th>
                <th>Type</th>
                <th>Market</th>
                <th>Added At</th>
                <th>Added By</th>
                <th>Actions</th>
            </tr>
            </thead>
            <tbody>
            {items.map((item) => (
                <tr key={item.id}>
                    <td>{item.name}</td>
                    <td>{item.qty}</td>
                    <td>{item.type}</td>
                    <td>{item.market}</td>
                    <td>{item.added_at}</td>
                    <td>{item.added_by}</td>
                    <td>
                        <button
                            className="btn btn-sm btn-danger"
                            onClick={() => handleDelete(item.id)}
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

export default ItemTable;