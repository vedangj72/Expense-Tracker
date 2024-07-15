import { createBrowserRouter, createRoutesFromElements, Route, RouterProvider } from 'react-router-dom'
import './App.css'
// import Login from './Components/Auth/Login'
import Navigation from './Components/Navigation/Navigation'
import Home from './Components/Home/Home'
import Signin from './Components/Auth/Signin'
import Login from './Components/Auth/Login'

function App() {
  const router=createBrowserRouter(
    createRoutesFromElements(
     <>
      <Route path='/' element={<Login/>}/>
      <Route path='/register' element={<Signin/>}/>
      <Route path='/layout' element={<Navigation/>}>
      <Route path='home' element={<Home/>}/>
      </Route>
     </>
    )
  )
  return (
    <>
    <RouterProvider router={router}></RouterProvider>
    </>
  )
}

export default App
