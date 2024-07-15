import { useState } from 'react';
import { useForm } from 'react-hook-form';
import axios from 'axios';
import { Link } from 'react-router-dom';
import { IoEyeSharp, IoEyeOffSharp } from 'react-icons/io5';

function Signin() {
    const { register, handleSubmit, formState: { errors } } = useForm();
    const [passwordVisibility, setPasswordVisibility] = useState(false);

    const onSubmit = async (data) => {
        try {
            const response = await axios.post('http://localhost:8080/register', data, {
                headers: {
                    'Content-Type': 'application/json',
                },
            });
            console.log('Success:', response.data);
            alert(`Registration successful! ${data.name}`);
        } catch (error) {
            alert("error in registration")
            console.error('Error:', error);
        }
    };

    const togglePasswordVisibility = () => {
        setPasswordVisibility(prevVisibility => !prevVisibility);
    };

    return (
        <div className="flex min-h-full flex-col justify-center px-6 py-12 lg:px-8 border-2 rounded-2xl " style={{backgroundColor:"#f3f4f6"}}>
            <div className="sm:mx-auto sm:w-full sm:max-w-sm">
                <h2 className="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900">Register your account</h2>
            </div>

            <div className="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
                <form className="space-y-6" onSubmit={handleSubmit(onSubmit)}>
                    <div className='text-left'>
                        <label htmlFor="name" className="block text-sm font-medium leading-6 text-gray-900">Name</label>
                        <div className="mt-2">
                            <input
                                id="name"
                                name="name"
                                type="text"
                                {...register('name', { required: 'Name is required' })}
                                className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                            />
                            {errors.name && <span className="text-red-500 text-sm">{errors.name.message}</span>}
                        </div>
                    </div>

                    <div className='text-left'>
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
                        <div className="mt-2">
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

                <p className="mt-10 text-center text-sm text-blue-500">
                    <Link to="/">Already have an account?</Link>
                </p>
            </div>
        </div>
    );
}

export default Signin;
