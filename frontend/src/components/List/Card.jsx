import React from 'react';
import './Card.scss'
import { Link } from 'react-router-dom';

const home = (props) => {
    const {data} = props

    return (
        <div className='outer-card'>
            <div className='card-border'>
                <div className='card-top'>
                    <Link to={`/getDish/${data.ID}`} >
                    <img src={data.link} alt={data.name} className='image'></img>
                    </Link>
                </div>
                <div className='card-bottom'>
                    <div className='data'>
                        <h2 className='rating'>
                        {data.name}
                        </h2>
                        <h4 className='rating'>{data.rating}</h4>
                        <h4 className='rating'>{data.price}</h4>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default home;