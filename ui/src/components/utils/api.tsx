export const getApiData = async (url:string, queryParams?: string) => {
    let apiUrl = import.meta.env.VITE_API_URL + "/"+ url
    if (queryParams != undefined) {
        apiUrl = apiUrl +   queryParams;
    }
    
    const token = localStorage.getItem("iggy_token");
    try {
        const response = await fetch(apiUrl, {
            method: 'GET',
             headers: {
                 'Content-Type': 'application/json',
                 'Authorization': 'Bearer ' + token,
             },
        });
        
        if (response.ok) {
            const result = await response.json();
            if (result.success) {
                return {"response": result.data} 
            } else {
                return {"error":result.data}
            }
        } 
        else {
            return {"error":response.statusText}
        }
    } catch (error) {
        return {"error":error}
    }
}



export function formatDate(isoString:string) {
    // Create a new Date object from the ISO string
    const date = new Date(isoString);
    
    // Extract the year, month, and day
    const year = date.getUTCFullYear();
    const month = String(date.getUTCMonth() + 1).padStart(2, '0');
    const day = String(date.getUTCDate()).padStart(2, '0');
    const hours = String(date.getUTCHours()).padStart(2, '0');
    const minutes = String(date.getUTCMinutes()).padStart(2, '0');
    const seconds = String(date.getUTCSeconds()).padStart(2, '0');
  
    
    // Return the formatted date as "YYYY-MM-DD"
    return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
  }


  export const sendApiData = async (url:string,postdata:any) => {
    const apiUrl = import.meta.env.VITE_API_URL + "/"+ url;
    const token = localStorage.getItem("iggy_token");
    try {
        const response = await fetch(apiUrl, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': 'Bearer ' + token,
            },
            body: JSON.stringify(postdata),
        });

        if (response.ok) {
            const result = await response.json();
            if (result.success) {
                return false
            } else {
                return { error: `Status: ${result.success} ${result.data}` };
            }
        } 
        else {
            return { error: `Error: ${response.status} ${response.statusText}` };
        }
    } catch (error) {
        console.error('Error:', error);
        return { error: (error as Error).message };
    }
}


export const sendApi = async (url:string,method:string,postdata:any) => {
    const apiUrl = import.meta.env.VITE_API_URL + "/"+ url;
    const token = localStorage.getItem("iggy_token");
    try {
        const response = await fetch(apiUrl, {
            method: method,
            headers: {
                'Content-Type': 'application/json',
                'Authorization': 'Bearer ' + token,
            },
            body: JSON.stringify(postdata),
        });

        if (response.ok) {
            const result = await response.json();
            if (result.success) {
                return false
            } else {
                return { error: `Status: ${result.success} ${result.data}` };
            }
        } 
        else {
            return { error: `Error: ${response.status} ${response.statusText}` };
        }
    } catch (error) {
        console.error('Error:', error);
        return { error: (error as Error).message };
    }
}



export const sendApi2 = async (url:string,method:string,postdata:any) => {
    const apiUrl = import.meta.env.VITE_API_URL + "/"+ url;
    const token = localStorage.getItem("iggy_token");
    try {
        const response = await fetch(apiUrl, {
            method: method,
            headers: {
                'Content-Type': 'application/json',
                'Authorization': 'Bearer ' + token,
            },
            body: JSON.stringify(postdata),
        });

        if (response.ok) {
            const result = await response.json();
            if (result.success) {
                return false
            } else {
                return { error: `Status: ${result.success} ${result.data}` };
            }
        } 
        else {
            return { error: `Error: ${response.status} ${response.statusText}` };
        }
    } catch (error) {
        console.error('Error:', error);
        return { error: (error as Error).message };
    }
}
