import React, { useEffect, useState } from 'react';
import './Dish.scss';
import dishApi from '../../common/apis/dishApi';
import { useParams } from 'react-router-dom';

const Dish = () => {
    const { id } = useParams();
    const [dishData, setDishData] = useState(null); 

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await dishApi.get(`/getDish/${id}`);
                setDishData(response.data); 
            } catch (error) {
                console.error(error);
            }
        };

        fetchData();
    }, [id]);

    return (
        <div className='Dish-outer'>
            <div className='Dish-back'>
                {dishData && (
                    <img src={dishData.link} alt={dishData.name} className='dish-image' />
                )}
            </div>
        </div>
    );
};

export default Dish;
