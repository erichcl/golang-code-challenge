import React from 'react';

export const Beer = ({product}) => {
    return <>
        <tr key={product.id} > 
            <td data-testid="td-name" width={150}>{product.name}</td>
            <td data-testid="td-temperature" width={150}>{product.temperature}</td>
            <td data-testid="td-temperatureStatus" width={150}>{product.temperatureStatus}</td>
        </tr>
    </>
}
