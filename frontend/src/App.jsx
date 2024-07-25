import {BrowserRouter as Router, Routes, Route} from 'react-router-dom'
import './App.scss'
import Home from './components/Home'
import Header from './components/Header/Header'
import Footer from './components/Footer/Footer'
import detail from './components/List/Dish'
import errorpage from './components/Error/Errorpage'

function App() {

  return (
   <div className='App'>
    <Router>
      <Header></Header>
      <div className='container'>
      <Routes>
        <Route path='/' Component={Home}></Route>
      <Route path='/movie/:imdbID' Component={detail}></Route>
      <Route Component={errorpage}></Route>
      </Routes>
      </div>
      <Footer></Footer>
    </Router>
   </div>
  )
}

export default App
