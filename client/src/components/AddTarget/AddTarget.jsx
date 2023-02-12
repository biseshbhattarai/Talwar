import axios from "axios";
import { Formik, Form } from "formik";
import { useNavigate } from "react-router-dom";
import Input from "../Input/Input";
import TopBar from "../TopBar/TopBar";
const AddTarget = () => {
    const navigate = useNavigate()

    const initialValues = {
        targetName: '',
        domain: ''
    }
    const handleSubmit = (values) => {

        const { targetName, domain } = values;
        const domains = []
        domain.split(',').forEach((domain) => {
            domains.push({
                Name: domain.trim(),
                IsScanned: false
            })
        })
        const data = {
            name: targetName,
            isScope: true,
            isScanned: true,
            domains: domains
        }
        axios.post('http://localhost:8000/api/targets', data)
            .then((response) => {
                if (response.status === 201) {
                    console.log("Added successfully")
                    navigate('/')
                }
            })
            .catch((error) => { console.log(error) })
    }
    return (
        <>
            <TopBar />
            <div className="bg-white">
                <div >
                    <div className="container mx-auto flex flex-wrap items-center">
                        <div className="lg:pr-32 md:py-32 w-full lg:w-1/2 flex flex-col bg-white">
                            <Formik
                                onSubmit={handleSubmit}
                                initialValues={initialValues}
                            >
                                <Form className="mt-11">
                                    <div className="flex gap-4">
                                        <Input name="targetName" id="targetName" label="Target Name" />
                                        <Input name="domain" id="domain" label="Domain" />
                                    </div>
                                    <div className="mt-6">
                                        <button
                                            className="py-3 text-white w-full bg-red-500 rounded font-semibold"
                                            type="submit"
                                        >
                                            Add
                                        </button>
                                    </div>
                                </Form>
                            </Formik>
                        </div>

                        <div className="pl-36 w-full lg:w-1/2">

                            <div className="mt-12 flex items-center">

                                <span className="ml-2 text-purple-800 text-4xl font-mono">
                                    Add New Target
                                </span>
                            </div>

                        </div>
                    </div>
                </div>
            </div>
        </>
    )
}

export default AddTarget;
