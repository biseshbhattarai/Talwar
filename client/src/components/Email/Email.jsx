import axios from "axios";
import { Formik, Form } from "formik";
import { useNavigate } from "react-router-dom";
import Input from "../Input/Input";
import Select from "../Select/Select"
import TopBar from "../TopBar/TopBar";

const Email = () => {
    const navigate = useNavigate()
    const handleSubmit = (values) => {
        axios.post('http://localhost:8000/api/targets/emails', {
            email: values.email,
            scanType: values.scanType
        })
            .then((response) => {
                if (response.status === 201) {
                    navigate('/')
                }
            })
            .catch((e) => {
                console.log(e)
            })
    }

    const initialValues = {
        email: '',
        scanType: '',
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
                                enableReinitialize
                            >
                                {({ setFieldValue, values }) => (
                                    <Form className="mt-11">
                                        <div className="flex flex-col gap-4">
                                            <Input name="email" id="email" label="Email" />
                                            <small className="text-xm -mt-4 italic">(Note : The above added email will receive update after the scan selected is completed)</small>
                                            <Select
                                                label="Scan Type"
                                                defaultValue={{ value: null, label: 'Select One' }}
                                                options={[
                                                    { value: 'subdomain', label: 'Subdomain' },
                                                    { value: 'githubscan', label: 'GithubScan' },
                                                    { value: 'portscan', label: 'PortScan' },
                                                    { value: 'linkedin', label: 'Linkedin' },
                                                    { value: 'nucleiscan', label: 'Nuclei Scan' },
                                                ]}
                                                onChange={(value) => setFieldValue('scanType', value)}
                                            />
                                        </div>
                                        <div className="mt-6">
                                            <button
                                                className="py-3 text-white w-full bg-red-500 rounded font-semibold"
                                                type="submit"
                                            >
                                                Save changes
                                            </button>
                                        </div>
                                    </Form>
                                )}
                            </Formik>
                        </div>

                        <div className="pl-36 w-full lg:w-1/2">

                            <div className="mt-12 flex items-center">

                                <span className="ml-2 text-purple-800 text-4xl font-mono">
                                    Email Settings
                                </span>
                            </div>

                        </div>
                    </div>
                </div>
            </div>
        </>
    )
}

export default Email;