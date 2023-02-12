import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import reportWebVitals from './reportWebVitals';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import Target from './components/Target/Target';
import Error from './components/ErrorPage/ErrorPage';
import AddTarget from './components/AddTarget/AddTarget';
import TargetDetail from './components/TargetDetail/TargetDetail';
import SubDomain from './components/SubDomain/SubDomain';
import ScanHistory from './components/ScanHistory/ScanHistory';
import ScanSchedule from './components/ScanSchedule/ScanSchedule';
import AddScheduleScan from './components/AddScheduleScan/AddScheduleScan';
import Vulnerabilities from './components/Vulnerabilities/Vulnerabilities';
import GitReport from './components/GitReport/GitReport';
import Email from './components/Email/Email';


const router = createBrowserRouter([
  {
    path: '/',
    element: <Target />,
    errorElement: <Error />
  },
  {
    path: '/add',
    element: <AddTarget />,
    errorElement: <Error />
  },
  {
    path: '/target/:id',
    element: <TargetDetail />,
    errorElement: <Error />
  },
  {
    path: '/target/:id/subdomains',
    element: <SubDomain />,
    errorElement: <Error />
  },
  {
    path: '/target/:id/history',
    element: <ScanHistory />,
    errorElement: <Error />
  },
  {
    path: '/schedule',
    element: <ScanSchedule />,
    errorElement: <Error />
  },
  {
    path: '/target/:id/addSchedule',
    element: <AddScheduleScan />,
    errorElement: <Error />
  },
  {
    path: '/target/:id/vulnerabilities',
    element: <Vulnerabilities />,
    errorElement: <Error />
  },
  {
    path: '/target/:id/scanReports',
    element: <GitReport />,
    errorElement: <Error />
  },
  {
    path: '/email-settings',
    element: <Email />,
    errorElement: <Error />
  }


])

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
