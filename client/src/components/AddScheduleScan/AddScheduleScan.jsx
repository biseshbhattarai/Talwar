import { Formik, Form } from "formik"
import Select from "../Select/Select"
import Input from "../Input/Input"
import { useNavigate, useParams } from "react-router-dom"
import { DatePicker } from "../DatePicker/DatePicker"
import axios from "axios"
import TopBar from "../TopBar/TopBar"

const AddScheduleScan = () => {
    const navigate = useNavigate()

    const { id } = useParams()


    const initialValues = {
        scanType: '',
        scanStart: '',
        scanInterval: ''
    }

    const handleSubmit = (values) => {
        const body = {
            scanType: values.scanType,
            scanStart: values.scanStart,
            scanInterval: +values.scanInterval,
        }
        // api/targets/12/scanSchedule
        axios.post('http://localhost:8000/api/targets/' + id + '/scanSchedule', body)
            .then((response) => {
                console.log(response)
                if (response.status === 201) {
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
                                enableReinitialize
                            >
                                {({ setFieldValue, values }) => (
                                    <Form className="mt-11">
                                        <div className="flex flex-col space-y-10">
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
                                            <Input name="scanInterval" id="scanInterval" label="Scan Interval" />
                                            <DatePicker
                                                selected={values.scanStart}
                                                label="Scan Start Date"
                                                name="scanStart"
                                                onChange={(value) => setFieldValue('scanStart', value)}
                                                dateFormat="dd/MM/yyyy"
                                                placeholderText="DD  /  MM  / YYYYY"
                                            />
                                        </div>
                                        <div className="mt-6">
                                            <button
                                                className="py-3 text-white w-full bg-red-500 rounded font-semibold"
                                                type="submit"
                                            >
                                                Create
                                            </button>
                                        </div>
                                    </Form>
                                )}


                            </Formik>
                        </div>

                        <div className="pl-36 w-full lg:w-1/2">

                            <div className="mt-12 flex items-center">

                                <span className="ml-2 text-purple-800 text-4xl font-mono">
                                    Schedule a Scan
                                </span>
                            </div>

                        </div>
                    </div>
                </div>
            </div>
        </>
    )

}

export default AddScheduleScan;