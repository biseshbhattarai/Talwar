import { useEffect, useState } from "react";
import axios from "axios";
import { useParams } from "react-router-dom";
import TopBar from "../TopBar/TopBar";
import { BiRefresh } from 'react-icons/bi'
import { toast, ToastContainer } from "react-toastify";



const Vulnerabilities = () => {
    const [vulnerabilities, setVulnerabilities] = useState([]);

    const { id } = useParams();

    const fetchVulnerabilities = () => {
        axios.get(`http://localhost:8000/api/targets/${id}/vulnerabilities`)
            .then((response) => {
                setVulnerabilities(response.data.nucleiScan);
            })
            .catch((error) => { console.log(error) })
    }

    useEffect(() => {
        fetchVulnerabilities();
    }, [])

    return (
        <>
            <TopBar />
            <h1 className="text-2xl text-center text-bold mt-2 mb-2">Vulnerabilities</h1>
            <div className='w-[700px] h-full bg-white border border-gray-ternary shadow rounded-md p-6 mt-4 mx-auto font-bold'>
                <div onClick={() => {
                    axios.get(`http://localhost:8000/api/targets/${id}/vulnerabilities`)
                        .then((response) => {
                            if (response.data.nucleiScan.length === 0) {
                                toast("Please wait . . .")
                            }
                            setVulnerabilities(response.data.nucleiScan);
                        })
                        .catch((error) => { console.log(error) })
                }}>
                    <BiRefresh />

                </div>
                {vulnerabilities?.length > 0 ? vulnerabilities.map((vulnerabilities, index) => (
                    <>
                        <div className="flex flex-col">
                            <p>Host : {vulnerabilities.Host}</p>
                            <p>Name : {vulnerabilities.Name}</p>
                            <p>Description : <small>{vulnerabilities.Description}</small></p>
                            <p className="text-blue-500">Severity : {vulnerabilities.Severity}</p>
                        </div>
                        <br />
                        <hr />
                    </>
                )) : (<p className="text-center">No Vulnerabilities yet!</p>)}

            </div>
            <ToastContainer />
        </>

    )
}

export default Vulnerabilities;