import axios from "axios";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import TopBar from "../TopBar/TopBar";
import { BiRefresh } from 'react-icons/bi'
import { toast, ToastContainer } from "react-toastify";


const SubDomain = () => {
    const { id } = useParams();

    const [subdomains, setSubdomains] = useState([])

    const fetchSubdomains = () => {
        axios.get('http://localhost:8000/api/targets/' + id + '/subdomains')
            .then((response) => {
                setSubdomains(response.data.subdomains)
            })
            .catch((error) => { console.log(error) })
    }



    useEffect(() => {
        fetchSubdomains()
    }
        , [])

    return (
        <>
            <TopBar />
            <br />
            <h1 className="text-2xl text-center text-bold mt-2 mb-2">Subdomains</h1>

            <div className='w-[700px] h-full bg-white border border-gray-ternary shadow rounded-md p-6 mt-2 mx-auto font-bold'>
                <input type="text" name="subdomainsearch" id="" placeholder="Search.." onChange={
                    (e) => {
                        const search = e.target.value
                        if (search != "") {
                            const filteredSubdomains = subdomains.filter((subdomain) => {
                                return subdomain.Name.includes(search)
                            })
                            setSubdomains(filteredSubdomains)
                        }

                    }
                } />
                <div className="mt-4 mb-1" onClick={() => {
                    axios.get('http://localhost:8000/api/targets/' + id + '/subdomains')
                        .then((response) => {
                            if (response.data.subdomains === null) {
                                toast("Please wait . . .")
                            }
                            setSubdomains(response.data.subdomains)
                        })
                        .catch((error) => { console.log(error) })
                }}>
                    <BiRefresh />

                </div>
                {subdomains?.length > 0 ? subdomains.map((subdomain, index) => (
                    <><a href={`https://${subdomain.Name}`}>{subdomain.Name}</a><br /></>
                )) : (<p className="text-center">No Subdomains yet!</p>)
                }
            </div>
            <ToastContainer />

        </>
    )

}
export default SubDomain;