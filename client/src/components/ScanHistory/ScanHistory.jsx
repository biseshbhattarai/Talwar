import axios from "axios";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import TopBar from "../TopBar/TopBar";

const ScanHistory = () => {
    const [scanHistory, setScanHistory] = useState([]);
    const { id } = useParams();

    useEffect(() => {
        axios.get('http://localhost:8000/api/targets/' + id + '/scanHistory')
            .then((response) => {
                setScanHistory(response.data.scan_history)
            })
            .catch((error) => { console.log(error) })
    }, [])
    return (
        <>
            <>
                <TopBar />
                <h1 className="text-2xl text-center text-bold mt-2 mb-2">Scan Histories</h1>

                {scanHistory?.length > 0 ? scanHistory.map((history, index) => (
                    <div key={index} className='w-[700px] h-[100px] bg-white border border-gray-ternary shadow rounded-md p-6 mt-2 mx-auto font-bold'>

                        <>
                            <p>ScanType : {history.ScanType}</p>
                            <p>Scanned At : {history.ScannedAt}</p>
                        </>

                    </div>
                )) : (<p className="text-center">No Scan ran yet!</p>)}


            </>
        </>
    )
}

export default ScanHistory;