import React, { useState } from "react";

const AddItemForm = ({ onAdd }) => {
    const [formData, setFormData] = useState({
        name: "",
        quantity: "",
        type: "",
        market: "",
        addedBy: "",
    });

    const handleChange = (e) => {
        setFormData({
            ...formData,
            [e.target.name]: e.target.value,
        });
    };

    const handleSubmit = (e) => {
        e.preventDefault();
        onAdd(formData);
        setFormData({ name: "", quantity: "", type: "", market: "", addedBy: "" });
    };

    return (
        <form onSubmit={handleSubmit}>
            <div className="row">
                <div className="col">
                    <input
                        type="text"
                        name="name"
                        className="form-control"
                        placeholder="Name"
                        value={formData.name}
                        onChange={handleChange}
                        required
                    />
                </div>
                <div className="col">
                    <select
                        name="quantity"
                        className="form-control"
                        value={formData.quantity}
                        onChange={handleChange}
                        required
                    >
                        <option value="">Select Quantity</option>
                        <option value="1">1</option>
                        <option value="2">2</option>
                        <option value="3">3</option>
                    </select>
                </div>
                <div className="col">
                    <select
                        name="type"
                        className="form-control"
                        value={formData.type}
                        onChange={handleChange}
                        required
                    >
                        <option value="">Select Type</option>
                        <option value="Food">Food</option>
                        <option value="Clothing">Clothing</option>
                    </select>
                </div>
                <div className="col">
                    <select
                        name="market"
                        className="form-control"
                        value={formData.market}
                        onChange={handleChange}
                        required
                    >
                        <option value="">Select Market</option>
                        <option value="Online">Online</option>
                        <option value="Retail">Retail</option>
                    </select>
                </div>
                <div className="col">
                    <input
                        type="text"
                        name="addedBy"
                        className="form-control"
                        placeholder="Added By"
                        value={formData.addedBy}
                        onChange={handleChange}
                        required
                    />
                </div>
                <div className="col">
                    <button type="submit" className="btn btn-sm btn-success">
                        Save
                    </button>
                </div>
            </div>
        </form>
    );
};

export default AddItemForm;