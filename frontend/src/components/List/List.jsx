import React from 'react';
import './List.scss'
import { useSelector } from 'react-redux';
import { getAllDish } from '../../redux/dishSlice';
import Card from './Card'

const List = () => {
    const dishes = useSelector(getAllDish)
    let render = dishes.map((dish, id) => {
        return <Card key = {id} data = {dish}></Card>
    });
    return (
        <div className = 'outer'>
            <div className ='list'>
                <div className='dish-container'>
                    {render}
                </div>
            </div>
        </div>
    );
};

export default List;