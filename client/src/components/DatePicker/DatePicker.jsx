import { useField } from 'formik';
import { useState } from 'react';
import ReactDatePicker, { ReactDatePickerProps } from 'react-datepicker';
import 'react-datepicker/dist/react-datepicker.css';
/*
This date component is an abstraction of the react date picker
it just adds label and some styles to match the design
passing any props to react date picker is just a matter of passing them to the component
*/



export const DatePicker = ({
    label,
    error,
    touched,
    name,
    disabled,
    ...props
}) => {
    function classNames(...classes) {
        return classes.filter(Boolean).join(' ')
    }
    const [, meta, helpers] = useField(name || 'date');
    return (
        <div>
            <label
                className="text-xs font-medium text-gray-800"
                htmlFor={props.id || name}
            >
                {label}
            </label>
            <div className="relative flex items-center">
                <ReactDatePicker
                    {...props}
                    className={classNames(
                        'mt-2 px-6 py-3 w-full rounded bg-white border border-gray-200',
                        'text-sm text-gray-900 font-medium focus: outline-info-400'
                    )}
                    autoComplete="off"
                    onFocus={() => {
                        return helpers.setTouched(true);
                    }}
                    useWeekdaysShort
                    disabled={disabled}
                    dateFormat={props.dateFormat ?? 'dd/MM/yyyy'}
                    strictParsing
                    id={props.id || name}
                />
                <label
                    className="absolute right-3 mt-2 text-gray-800"
                    htmlFor={props.id || name}
                >
                </label>
            </div>
            {error && meta.touched && (
                <span className="text-xs text-error-500 mt-1" role="alert">
                    {error}
                </span>
            )}
        </div>
    );
};
