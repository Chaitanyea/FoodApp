import React from 'react';
import { useState } from 'react';
import dishApi from '../../common/apis/dishApi';
import './addDishPage.scss'
import { Navigate, useNavigate } from 'react-router';

const addDishPage = () => {
    const [name, setName] = useState("");
    const [price, setPrice] = useState(0);
    const [rating, setRating] = useState(0);
    const [link, setLink] = useState("");
    const navigate = useNavigate();

    const handleSave = () => {
        const data = {
            name,
            price: parseInt(price), 
            rating : parseFloat(rating),
            link
        }
        dishApi.put('/createDish', data)
        .then(() => {
            console.log(data)
            navigate('/');
        })
        .catch((err) => {
            console.log(err);
            //navigate('/error/');
        });
    }

    return (
        <div className='add-outer'>
            <div className='add-border'>
                <div className='add-coloumn'>
                    <label className='add-heading'>Name</label>
                    <input 
                        type='text'
                        value={name}
                        onChange={(e) => setName(e.target.value)}
                        className='input-tab'
                    />
                </div>

                <div className='add-coloumn'>
                    <label className='add-heading'>Price</label>
                    <input 
                        type='text'
                        value={price}
                        onChange={(e) => setPrice(e.target.value)}
                        className='input-tab'
                    />
                </div>

                <div className='add-coloumn'>
                    <label className='add-heading'>Rating</label>
                    <input 
                        type='text'
                        value={rating}
                        onChange={(e) => setRating(e.target.value)}
                        className='input-tab'
                    />
                </div>

                <div className='add-coloumn'>
                    <label className='add-heading'>Link</label>
                    <input 
                        type='text'
                        value={link}
                        onChange={(e) => setLink(e.target.value)}
                        className='input-tab'
                    />
                </div>

                <button className='save-button' onClick={handleSave}>
                    <h1 className='text-save'>
                        Save
                    </h1>
                </button>
            </div>
        </div>
    );
};

export default addDishPage;