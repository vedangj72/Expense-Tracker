import { useState } from 'react';
import { useForm } from 'react-hook-form';
import axios from 'axios';
import { Link, useNavigate } from 'react-router-dom';
import { IoEyeOffSharp, IoEyeSharp } from 'react-icons/io5';

function LoginForm() {
    const { register, handleSubmit, formState: { errors } } = useForm();
    const [passwordVisibility, setPasswordVisibility] = useState(false);
    const navigate=useNavigate();

    const onSubmit = async (data) => {
        try {
            const response = await axios.post('http://localhost:8080/login', data, {
                headers: {
                    'Content-Type': 'application/json',
                },
            });
            console.log('Success:', response.data);

            // console.log(response.data)
            // Store JWT token in localStorage
            localStorage.setItem('token', response.data.token);

            // Redirect to home page or any other route after successful login
           navigate('/layout/home');
           
        } catch (error) {
          alert("Wrong credentials");
            console.error('Error:', error);
        }

    };
    const togglePasswordVisibility = () => {
      setPasswordVisibility(prevVisibility => !prevVisibility);
  };
    return (
        <div className="flex min-h-full flex-col justify-center px-6 py-12 lg:px-8 border-2 rounded-2xl"style={{backgroundColor:"#f3f4f6"}}>
            <div className="sm:mx-auto sm:w-full sm:max-w-sm">
                <h2 className="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900">Log in to your Account</h2>
            </div>

            <div className="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
                <form className="space-y-6" onSubmit={handleSubmit(onSubmit)}>

                    <div className=' text-left'>
                        <label htmlFor="email" className="block text-sm font-medium leading-6 text-gray-900">Email address</label>
                        <div className="mt-2">
                            <input
                                id="email"
                                name="email"
                                type="email"
                                autoComplete="email"
                                {...register('email', { required: 'Email is required' })}
                                className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                            />
                            {errors.email && <span className="text-red-500 text-sm">{errors.email.message}</span>}
                        </div>
                    </div>

                    <div>
                        <div className="flex items-center justify-between">
                            <label htmlFor="password" className="block text-sm font-medium leading-6 text-gray-900">Password</label>
                            <button type="button" onClick={togglePasswordVisibility} className="text-gray-500 hover:text-gray-700 focus:outline-none">
                                {passwordVisibility ? <IoEyeOffSharp /> : <IoEyeSharp />}
                            </button>
                        </div>
                        <div className="mt-2 text-right" >
                      
                            <input
                                id="password"
                                name="password"
                                type={passwordVisibility ? 'text' : 'password'}
                                autoComplete="current-password"
                                {...register('password', { required: 'Password is required' })}
                                className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                            />
                            {errors.password && <span className="text-red-500 text-sm">{errors.password.message}</span>}
                        </div>
                    </div>

                    <div>
                        <button type="submit" className="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">Sign in</button>
                    </div>
                </form>

                <p className="mt-10 text-center text-sm text-blue-500 ">
                    <Link to="/register">Do not have an account? Register now</Link>
                </p>
            </div>
        </div>
    );
}

export default LoginForm;
