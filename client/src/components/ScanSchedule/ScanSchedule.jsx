import axios from "axios";
import { useEffect, useState } from "react";
import TopBar from "../TopBar/TopBar";

const ScanSchedule = () => {
    const [schedule, setSchedule] = useState([])

    useEffect(() => {
        axios.get('http://localhost:8000/api/targets/scanSchedule')
            .then((response) => {
                setSchedule(response.data.scanSchedules)
            })
            .catch((error) => { console.log(error) })
    }, [])

    return (
        <>
            <>
                <TopBar />
                <h1 className="text-2xl text-center text-bold mt-2 mb-2">Scan Schedules</h1>

                {schedule.length ? schedule.map((s, index) => (
                    <div key={index} className='w-[700px] h-[150px] bg-white border border-gray-ternary shadow rounded-md p-6 mt-2 mx-auto font-bold '>

                        <>
                            <p className="text-xl">Target Name - {s.Target.Name}</p>
                            <p>ScanType : {s.ScanType}</p>
                            <p>Scanned Start : {s.ScanStart}</p>
                            <p>Scan Interval : {s.ScanInterval} min</p>

                        </>

                    </div>
                )) : <p className="text-center text-xl mt-2">No Scan Schedule yet !</p>}


            </>
        </>
    )
}

export default ScanSchedule;