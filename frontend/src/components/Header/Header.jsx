import React from 'react';
import {Link} from 'react-router-dom'
import './Header.scss'
import wine from '../../assets/wine.png'

const header = () => {
    return (
        <div className='header'>
            <Link to= "/" className='Link'>
                <div className='logo'> Brominos </div>
            </Link>
            <Link to ='/addDish/'>
                <div className='wine-image'>
                <img src= {wine} alt='wine'></img>
            </div>
            </Link>
        </div>
    )
};

export default header;