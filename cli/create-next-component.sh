#!/bin/bash

# directory name passed as a parameter
directory_name=$1

# function to convert this-case to PascalCase
convert_to_pascal() {
    local input="$1"
    local output=""
    IFS='-' read -ra words <<< "$input"
    for word in "${words[@]}"; do
        output+="${word^}"
    done
    echo "$output"
}

# make the directory for the component
mkdir -p ./${directory_name}
# create the .css file
touch ./${directory_name}/${directory_name}.css
# create the .tsx file
component_name_pascal=$(convert_to_pascal ${directory_name})
component_file=./${directory_name}/${component_name_pascal}.tsx
touch ${component_file}

echo "import './${directory_name}.css'

export default function ${component_name_pascal}() {
    return (
        <>

        </>
    )
}
" > ${component_file}