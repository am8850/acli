import { useEffect, useState } from "react"
import { GetDataAsync } from "./services"

export interface IWeather {
    city: string
    temperature: number
    condition: string
}

const App = () => {
    const [weather, setWeather] = useState<IWeather[]>([])

    const LoadData = async () => {
        let data = (await GetDataAsync("/api/weather")) as IWeather[]
        setWeather(data)
    }

    useEffect(() => {
        LoadData()
    }, [])

    return (
        <div className="container mx-auto">
            <h1>Vite+Go+TailwindCSS</h1>
            {weather.map((w, i) => <ul>
                <li key={i}>{w.city} - {w.temperature} - {w.condition}</li>
            </ul>)}
        </div>
    )
}

export default App