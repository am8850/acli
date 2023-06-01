import axios from 'axios'
import { ICompletion, IPrompt } from '../interfaces'

const config = {
    headers: {
        'Content-Type': 'application/json',
        'api-key': import.meta.env.VITE_OPENAI_KEY,
    },
}

export async function OpenAIServiceAsync(
    model: string,
    payload: IPrompt
): Promise<ICompletion> {
    let url = ''
    if (model === 'Davinci') {
        url = import.meta.env.VITE_OPENAI_DAVINCI_URL
        delete payload.messages
    } else {
        url = import.meta.env.VITE_OPENAI_GTP_URL
        delete payload.prompt
    }
    console.info('Request:', payload)
    const req = await axios.post(url, payload, config)
    console.info('Response:', req.data)
    return req.data as ICompletion
}

export function OpenAIServicePostPromise(
    model: string,
    payload: IPrompt
): Promise<ICompletion> {
    let url = ''
    if (model === 'Davinci') {
        url = import.meta.env.VITE_OPENAI_DAVINCI_URL
        delete payload.messages
    } else {
        url = import.meta.env.VITE_OPENAI_GTP_URL
        delete payload.prompt
    }
    console.info('Request:', payload)
    return axios.post(url, payload, config)
}

export async function PromiseAllAsync(promises: any[]): Promise<any[]> {
    let responses = await axios.all(promises)
    return responses
}

