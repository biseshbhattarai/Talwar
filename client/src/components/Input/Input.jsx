import { FieldHookConfig, useField } from 'formik';
import React from 'react';


const Input = ({ label, ...props }) => {
    const [field, meta] = useField(props);
    return (
        <div className="flex flex-col w-full">
            {label && (
                <label
                    className="text-xs font-medium text-gray-800"
                    htmlFor={props.id || props.name}
                >
                    {label}
                </label>
            )}
            <input
                className="mt-1 px-4 py-3 border border-gray-100 rounded focus:outline-info-400"
                {...field}
                id={props.id || props.name}
                disabled={props.disabled}
            />
            {meta.touched && meta.error ? (
                <span className="text-xs text-error-500 mt-1" role="alert">
                    {meta.error}
                </span>
            ) : null}
        </div>
    );
};

export default Input;
