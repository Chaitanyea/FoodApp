import React, { useEffect, useState } from 'react';
import './Dish.scss';
import dishApi from '../../common/apis/dishApi';
import { Link, useNavigate, useParams } from 'react-router-dom';
import Icon from 'react-crud-icons';

const Dish = () => {
    const { id } = useParams();
    const [dishData, setDishData] = useState(null); 

    //UpdateDish constants
    const [name, setName] = useState("");
    const [price, setPrice] = useState(0);
    const [rating, setRating] = useState(0);
    const [link, setLink] = useState("");
    const navigate = useNavigate();

    const handleSave = async () => {
        try {
            await dishApi.delete(`/deleteDish/${id}`)
        } catch (error) {
            console.error(error);
            useNavigate('/error/')
        }
        const data = {
            name,
            price: parseInt(price),
            rating: parseFloat(rating),
            link
        }
        await dishApi.put(`/createDish`, data)
        .then(() => {
            navigate('/')
        })
        .catch((err) => {
            console.log(err);
            navigate('/error/');
        });
    }

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await dishApi.get(`/getDish/${id}`);
                console.log(response)
                setDishData(response.data); 
                setLink(response.data.link);
                setName(response.data.name);
                setPrice(response.data.price);
                setRating(response.data.rating);
            } catch (error) {
                console.error(error);
            }
        };

        fetchData();
    }, [id]);

    const handleDelete = async () => {
        try {
            await dishApi.delete(`/deleteDish/${id}`)
            .then(() => {
                useNavigate('/')
            });
        } catch (error) {
            console.error(error);
            useNavigate('/error/')
        }
    };

    return (
        <div className='Dish-outer'>
            <div className='Dish-back'>
                {dishData && (
                    <img src={dishData.link} alt={dishData.name} className='dish-image' />
                )}
                <div className='icon'> 
                    <Link to= '/'>
                        <Icon
                        name="delete"
                        theme="light"
                        size="small"
                        onClick={handleDelete}
                    />
                    </Link>
                </div>
            </div>
            {dishData && (<div className='add-border'>
                <div className='add-coloumn'>
                    <label className='add-heading'>Name</label>
                    <input 
                        type='text'
                        onChange={(e) => setName(e.target.value)}
                        className='input-tab'
                        defaultValue={dishData.name}
                    />
                </div>

                <div className='add-coloumn'>
                    <label className='add-heading'>Price</label>
                    <input 
                        type='text'
                        onChange={(e) => setPrice(e.target.value)}
                        className='input-tab'
                        defaultValue={dishData.price}
                    /> 
                </div>

                <div className='add-coloumn'>
                    <label className='add-heading'>Rating</label>
                    <input 
                        type='text'
                        onChange={(e) => setRating(e.target.value)}
                        className='input-tab'
                        defaultValue={dishData.rating}
                    />
                </div>

                <div className='add-coloumn'>
                    <label className='add-heading'>Link</label>
                    <input 
                        type='text'
                        onChange={(e) => setLink(e.target.value)}
                        className='input-tab'
                        defaultValue={dishData.link}
                    />
                </div>

                <button className='save-button' onClick={handleSave}>
                    <h1 className='text-save'>
                        Save
                    </h1>
                </button>
            </div>)}
        </div>
    );
};

export default Dish;
