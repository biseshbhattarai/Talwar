import TopBar from "../TopBar/TopBar";
import { useParams } from "react-router-dom";
import { useEffect, useState } from "react";
import axios from "axios";
import Input from "../Input/Input";
import { BiRefresh } from 'react-icons/bi'
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

const TargetDetail = () => {
    const { id } = useParams();
    const [subdomains, setSubdomains] = useState([])
    const [ports, setPorts] = useState([])
    const [repos, setRepos] = useState([])
    const [organization, setOrganization] = useState([])
    const [organizationName, setOrganizationName] = useState('')

    const fetchSubdomains = () => {
        axios.get('http://localhost:8000/api/targets/' + id + '/subdomains')
            .then((response) => {
                setSubdomains(response.data.subdomains)
            })
            .catch((error) => { console.log(error) })
    }

    const fetchOrganizationRepos = () => {
        axios.get('http://localhost:8000/api/targets/' + id + '/organization/repos')
            .then((response) => {
                setRepos(response.data.organizationRepos)
            })
            .catch((error) => { console.log(error) })
    }

    const fetchPorts = () => {
        axios.get('http://localhost:8000/api/targets/' + id + '/ports')
            .then((response) => {
                setPorts(response.data.portscan)
            })
            .catch((error) => { console.log(error) })
    }

    const fetchOrganization = () => {
        axios.get('http://localhost:8000/api/targets/' + id + '/organization')
            .then((response) => {
                setOrganization(response.data.organizationRepos)
            })
            .catch((error) => { console.log(error) })
    }

    useEffect(() => {
        fetchSubdomains()
        fetchOrganizationRepos()
        fetchPorts()
        fetchOrganization()
    }, [])

    const submitOrg = () => {
        axios.post('http://localhost:8000/api/targets/' + id + '/organization', {
            "OrgName": organizationName,
            "OrganizationType": "github"
        })
            .then((response) => {
                console.log(response)
                fetchOrganization()
            })
            .catch((error) => { console.log(error) })
    }

    const searchGithubRepos = () => {
        axios.post('http://localhost:8000/api/targets/' + id + '/organization/startSearchRepo', {})
            .then((response) => {
                console.log(response)
                toast("Github repos scanning started")
            })
            .catch((error) => { console.log(error) })
    }

    const searchSubDomains = () => {
        axios.post('http://localhost:8000/api/targets/' + id + '/startScan')
            .then((response) => {
                console.log(response)
                toast("Subdomain scanning started")

            })
            .catch((error) => { console.log(error) })
    }

    const searchPorts = () => {
        axios.post('http://localhost:8000/api/targets/' + id + '/portScan', {})
            .then((response) => {
                console.log(response)
                toast("Ports scanning started")
            })
            .catch((error) => { console.log(error) })
    }

    const scanVulnerabilities = () => {
        axios.post('http://localhost:8000/api/targets/' + id + '/nucleiScan', {})
            .then((response) => {
                console.log(response.data)
                toast("Vulnerabilities scanning started")

            })
            .catch((error) => { console.log(error) })
    }

    const scanGithubRepos = () => {
        axios.post('http://localhost:8000/api/targets/' + id + '/startGithubScan', {})
            .then((response) => {
                console.log(response)
                toast("Github repo scanning started")

            })
            .catch((error) => { console.log(error) })
    }

    return (
        <>
            <TopBar />
            <div className="flex flex-row space-x-8 mt-4">
                <button
                    className="py-3 text-white w-full bg-slate-500 rounded font-semibold"
                    type="button"
                    onClick={searchSubDomains}
                >
                    Scan Subdomains
                </button>
                {
                    organization.length > 0 && (
                        <button
                            onClick={searchGithubRepos}
                            className="py-3 text-white w-full bg-slate-500 rounded font-semibold"
                            type="button"
                        >
                            Search github repos
                        </button>
                    )
                }

                {
                    subdomains?.length > 0 && (
                        <button
                            className="py-3 text-white w-full bg-slate-500 rounded font-semibold"
                            type="button"
                            onClick={searchPorts}
                        >
                            Scan ports
                        </button>
                    )
                }

                {
                    repos?.length > 0 && (
                        <button
                            className="py-3 text-white w-full bg-slate-500 rounded font-semibold"
                            type="button"
                            onClick={scanGithubRepos}
                        >
                            Scan github repos
                        </button>
                    )
                }
                {subdomains?.length > 0 && (

                    <button
                        className="py-3 text-white w-full bg-slate-500 rounded font-semibold"
                        type="submit"
                        onClick={scanVulnerabilities}
                    >
                        Scan vulnerabilities
                    </button>
                )}
            </div>
            <div className="flex flex-row space-x-8 mt-4">
                <div className='w-[900px] h-full bg-white border border-gray-ternary shadow rounded-md p-6 mt-2 mx-2 font-bold'>
                    <div className="flex justify-between items-start">
                        <h4>Subdomains</h4>
                        <div onClick={() => {
                            axios.get('http://localhost:8000/api/targets/' + id + '/subdomains')
                                .then((response) => {
                                    if (response.data.subdomains === null) {
                                        toast("Please wait ...")
                                    }
                                    setSubdomains(response.data.subdomains)
                                })
                                .catch((error) => { console.log(error) })
                        }}>
                            <BiRefresh />

                        </div>
                    </div>
                    <input type="text" name="subdomainsearch" id="" placeholder="Search.." onChange={
                        (e) => {
                            const search = e.target.value
                            console.log(search)
                            if (search != "") {
                                const filteredSubdomains = subdomains.filter((subdomain) => {
                                    return subdomain.Name.includes(search)
                                })
                                setSubdomains(filteredSubdomains)
                            }

                        }
                    } />
                    <br />
                    {subdomains?.length > 0 && subdomains.map((subdomain, index) => (
                        <><a href={`https://${subdomain.Name}`} _blank><small className="" key={index}>{subdomain.Name}</small></a><br /></>
                    )
                    )}
                </div>
                <div className='w-[700px] h-full bg-white border border-gray-ternary shadow rounded-md p-6 mt-2 mx-10 font-bold'>
                    <div className="flex justify-between items-start">
                        <h4 className="underline">Repositories</h4>
                        <div onClick={fetchOrganizationRepos}>
                            <BiRefresh />

                        </div>

                    </div>
                    {organization.length === 0 && (
                        <div>
                            <input type="text" placeholder="Enter organization name" name="organization" onChange={(e) => { setOrganizationName(e.target.value) }} />
                            <button onClick={submitOrg} className="py-2 text-white w-20 bg-slate-500 rounded font-semibold" type="submit">Submit</button>
                        </div>
                    )}
                    {repos?.length > 0 && repos.map((repo, index) => (
                        <><p key={index}>{repo.RepoName}</p><a className="text-blue-600" href={repo.RepoUrl}>{repo.RepoUrl}</a></>
                    )
                    )}
                </div>
                <div className='w-[700px] h-full bg-white border border-gray-ternary shadow rounded-md p-6 mt-2 mx-10 font-bold'>
                    <div className="flex justify-between items-start">
                        <h4 className="underline">Ports</h4>
                        <div onClick={() => {
                            axios.get('http://localhost:8000/api/targets/' + id + '/ports')
                                .then((response) => {
                                    if (response.data.portscan === null) {
                                        toast("Please wait ...")
                                    }
                                    setPorts(response.data.portscan)
                                })
                                .catch((error) => { console.log(error) })
                        }}>
                            <BiRefresh />

                        </div>

                    </div>
                    {ports?.length > 0 && ports.map((port, index) => (
                        <><small><a href={port.PortNumber} key={index}>{port.PortNumber}</a></small><br /></>
                    ))}
                </div>
            </div>
            <ToastContainer />

        </>
    )
}

export default TargetDetail;