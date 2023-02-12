import { useEffect, useState } from "react";
import axios from "axios";
import { useParams } from "react-router-dom";
import TopBar from "../TopBar/TopBar";
import { BiRefresh } from 'react-icons/bi'
import { ToastContainer, toast } from "react-toastify";


const GitReport = () => {
    const [gitReports, setGitReports] = useState([]);

    const { id } = useParams();

    const fetchGitReports = () => {
        axios.get(`http://localhost:8000/api/targets/${id}/scanResults`)
            .then((response) => {
                setGitReports(response.data.truffleScan);
            })
            .catch((error) => { console.log(error) })
    }

    useEffect(() => {
        fetchGitReports();
    }, [])
    return (
        <>
            <TopBar />
            <h1 className="text-2xl text-center text-bold mt-2 mb-2">Git Reports</h1>
            <div className='w-[700px] h-full bg-white border border-gray-ternary shadow rounded-md p-6 mt-2 font-bold mx-auto'>
                <div onClick={() => {
                    axios.get(`http://localhost:8000/api/targets/${id}/scanResults`)
                        .then((response) => {
                            if (response.data.truffleScan.length <= 0) {
                                toast("Please wait . . .")
                            }
                            setGitReports(response.data.truffleScan);
                        })
                        .catch((error) => { console.log(error) })
                }}>
                    <BiRefresh />

                </div>
                {gitReports?.length > 0 ? gitReports.map((gitReport, index) => (
                    <><div className="flex flex-col">
                        <p>Repository : {gitReport.Repository}</p>
                        <p>Commit : {gitReport.Commit}</p>
                        <p>File : <small>{gitReport.File}</small></p>
                        <p>Email : {gitReport.Email}</p>
                        <p>Redacted : {gitReport.Redacted}</p>
                    </div><br /><hr /></>
                )) : (<p className="text-center">No Git Reports yet!</p>)}

            </div>
            <ToastContainer />
        </>

    )
}

export default GitReport;