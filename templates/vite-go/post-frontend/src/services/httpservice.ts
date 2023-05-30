import axios from 'axios'

export async function GetDataAsync(url: string) {
    let response = await axios.get(url)
    return response.data
}