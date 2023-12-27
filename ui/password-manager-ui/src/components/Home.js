import {useState} from 'react'

const Home = () => {

    const link  = "http://localhost:8080/health"

    const [healthData, setHealthData] = useState("")
    const [errorData, setErrorData] = useState(null)
    const fetchData = () => {
        fetch(link)
        .then(res => res.json())
        .then(data => {
            console.log(data)
            setHealthData(data)
            setErrorData(null)
        })
        .catch(err => {
            console.log(err)
            setErrorData(err.message)
        })
    }

    fetchData()

    return (
        <div className="home">
            <h1>Password Manager!</h1>

            { healthData && <div className="health-data">{ healthData }</div> }
            { errorData && <div className="error-data">{ errorData }</div> }
        </div>
    );
}

export default Home