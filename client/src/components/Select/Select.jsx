import { Listbox, Transition } from '@headlessui/react';
import { Fragment, useEffect, useState } from 'react';



const DEFAULT_VALUE = {
    label: 'Select One',
    value: null,
};

const Select = ({
    label,
    options,
    onChange,
    optional,
    defaultValue,
}) => {

    function classNames(...classes) {
        return classes.filter(Boolean).join(' ')
    }

    const [selected, setSelected] = useState(defaultValue || DEFAULT_VALUE);

    // useEffect(() => {
    //     setSelected(defaultValue);
    // }, [defaultValue]);

    const handleChange = (value) => {
        setSelected(value);
        if (onChange) {
            onChange(value.value);

        }
    };
    return (
        <div className="flex flex-col">
            {label && (
                <label className="text-xs font-medium text-gray-800 flex justify-between items-center">
                    <span>{label}</span>
                    {optional && <span>Optional</span>}
                </label>
            )}
            <Listbox value={selected} onChange={handleChange}>
                <div className="mt-1 relative">
                    <Listbox.Button
                        className={classNames(
                            'relative w-full py-3 px-4 flex items-center text-dark bg-white rounded-lg',
                            'cursor-pointer focus:outline-none focus-visible:ring-2 focus-visible:ring-opacity-75 focus-visible:ring-white focus-visible:ring-offset-orange-300 focus-visible:ring-offset-2',
                            'focus-visible:border-indigo-500 sm:text-sm border border-gray-100'
                        )}
                    >
                        <span className="block font-medium">{selected?.label}</span>

                    </Listbox.Button>
                    <Transition
                        as={Fragment}
                        leave="transition ease-in duration-100"
                        leaveFrom="opacity-100"
                        leaveTo="opacity-0"
                    >
                        <Listbox.Options className="absolute z-10 w-full p-1 mt-1 overflow-auto text-base bg-white rounded-md shadow-lg max-h-60 ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm">
                            {options?.map((option, idx) => (
                                <Listbox.Option
                                    key={idx}
                                    className={({ active }) =>
                                        `${active ? 'text-black bg-gray-50' : 'text-gray-900'}
                         rounded cursor-pointer select-none py-2 px-3`
                                    }
                                    value={option}
                                >
                                    {({ selected, active }) => (
                                        <>
                                            <span
                                                className={`${selected ? 'font-medium' : 'font-normal'
                                                    } block`}
                                            >
                                                {option.label}
                                            </span>
                                            {selected ? (
                                                <span
                                                    className={classNames(
                                                        'absolute inset-y-0 left-0 flex items-center pl-3',
                                                        active ? 'text-amber-600' : 'text-amber-600'
                                                    )}
                                                ></span>
                                            ) : null}
                                        </>
                                    )}
                                </Listbox.Option>
                            ))}
                        </Listbox.Options>
                    </Transition>
                </div>
            </Listbox>
        </div>
    );
};

export default Select;
