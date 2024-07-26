import React, { useEffect } from 'react';
import wine from '../assets/wine.png'
import List from './List/List'
import './Home.scss'
import { useDispatch } from 'react-redux';
import { addDish } from '../redux/dishSlice';
import dishApi from '../common/apis/dishApi'

const Home = () => {
    const dispatch = useDispatch();
    useEffect(() =>{
        const getDish = async () => {
            const response = await dishApi.get('/getDish/')
            .catch((err) => {
                console.log(err)
            });
            dispatch(addDish(response.data))
        }
        getDish();
    }, [])  

    return (
        <div className='outer'>
            <div className='padding'>

            </div>
            <div className='list'>
                <img src={wine} ></img>
                <List>

                </List>
            </div>
        </div>
    );
};

export default Home;