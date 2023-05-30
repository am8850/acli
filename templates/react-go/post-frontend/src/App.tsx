import { useEffect, useState } from "react"
import { GetDataAsync } from "./services"

export interface IWeather {
    city: string
    temperature: number
    conditions: string
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
            <h1 className="font-bold text-lg">Vite+Go+TailwindCSS</h1>
            <br /><br />
            <label className="uppercase">Temperatures:</label>
            <br />
            {weather.map((w, i) => <ul>
                <li key={i}><span className="font-bold">{w.city}:</span> {w.temperature} - {w.conditions}</li>
            </ul>)}
        </div>
    )
}