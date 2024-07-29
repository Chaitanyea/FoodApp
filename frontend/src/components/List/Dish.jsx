import React, { useEffect, useState } from 'react';
import './Dish.scss';
import dishApi from '../../common/apis/dishApi';
import { Link, useNavigate, useParams } from 'react-router-dom';
import Icon from 'react-crud-icons';

const Dish = () => {
    const { id } = useParams();
    const [dishData, setDishData] = useState(null); 

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await dishApi.get(`/getDish/${id}`);
                console.log(response)
                setDishData(response.data); 
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
        </div>
    );
};

export default Dish;
