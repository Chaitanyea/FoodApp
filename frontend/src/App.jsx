import {BrowserRouter as Router, Routes, Route} from 'react-router-dom'
import './App.scss'
import Home from './components/Home'
import Header from './components/Header/Header'
import Footer from './components/Footer/Footer'
import detail from './components/List/Dish'
import errorpage from './components/Error/Errorpage'
import addDishPage from './components/AddDish/addDishPage'

function App() {

  return (
   <div className='App'>
    <Router>
      <Header></Header>
      <div className='container'>
      <Routes>
        <Route path='/' Component={Home}></Route>
      <Route path='/getDish/:id' Component={detail}></Route>
      <Route path ='/error/'Component={errorpage}></Route>
      <Route path ='/addDish/' Component = {addDishPage}></Route>
      </Routes>
      </div>
      <Footer></Footer>
    </Router>
   </div>
  )
}

export default App
