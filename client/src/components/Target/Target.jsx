import { useEffect, useState } from "react";
import axios from 'axios'
import TopBar from "../TopBar/TopBar";
import { FaArrowCircleRight } from 'react-icons/fa'
import { Link } from "react-router-dom";
import { useNavigate } from "react-router-dom";


const Target = () => {
    const [targets, setTargets] = useState([]);
    const navigate = useNavigate()


    useEffect(() => {
        axios.get('http://localhost:8000/api/targets')
            .then((response) => {
                setTargets(response.data.targets);
            })
            .catch((error) => {
                console.log(error)
            })
    }, [])

    const handleSearch = (e) => {
        const search = e.target.value;
        //local search 
        if (search === '') {
            setTargets(targets);  // reset targets to original value
            return;
        }
        const filtered = targets.filter((target) => {
            return target.Name.toLowerCase().includes(search.toLowerCase())
        })
        setTargets(filtered)
    }

    return (
        <>
            <TopBar />
            <div className="flex items-center justify-center mb-10 mt-4">
                <input className="border-b-2 w-[900px] text-center outline-none" type="text" placeholder="Search targets....." onChange={(e) => { handleSearch(e) }} />
            </div>
            {
                targets.length ? (
                    targets.map((target, index) => {
                        return (
                            <div key={index} className='w-[900px] h-[120px] bg-white border border-gray-ternary shadow rounded-md p-6 mt-2  font-bold mx-auto'>
                                <div className="flex justify-between items-start">
                                    <p>{target.Name}</p>
                                    <Link to={{ pathname: `/target/${target.ID}` }}>
                                        <FaArrowCircleRight />
                                    </Link>
                                </div>
                                <hr />
                                <div className="flex flex-row text-sm space-x-4 mt-2">
                                    <button className="p-auto text-white w-full h-8 bg-slate-500 rounded font-semibold" onClick={() => { navigate(`/target/${target.ID}/subdomains`) }}>
                                        View subdomains
                                    </button>
                                    <button className="p-auto text-white w-full h-8 bg-slate-500 rounded font-semibold" onClick={() => { navigate(`/target/${target.ID}/scanReports`) }}>
                                        View github reports
                                    </button>
                                    <button className="p-auto text-white w-full h-8 bg-slate-500 rounded font-semibold" onClick={() => { navigate(`/target/${target.ID}/vulnerabilities`) }}>
                                        View vulnerabilities
                                    </button>
                                    <button className="p-auto text-white w-full h-8 bg-slate-500 rounded font-semibold" onClick={() => { navigate(`/target/${target.ID}/history`) }}>
                                        View scan histories
                                    </button>
                                    <button className="p-auto text-white w-full h-8 bg-slate-500 rounded font-semibold" onClick={() => { navigate(`/target/${target.ID}/addSchedule`) }}>
                                        Add schedule scan
                                    </button>
                                </div>

                            </div>

                        )
                    }
                    )
                ) : (<p className="text-center">No targets yet!</p>)}
        </>
    )
}

export default Target; 
