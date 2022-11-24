type Props = {
  placeholder: string;
  value: string;
  onChange: (event: React.ChangeEvent<HTMLInputElement>) => void;
  onKeyDown?: (event: React.KeyboardEvent<HTMLInputElement>) => void | undefined;
};

function TextInput(props: Props) {
  return (
    <input
      className="py-4 px-3.5 rounded-2xl shadow appearance-none border leading-tight focus:outline-none focus:shadow-outline"
      value={props.value}
      onChange={props.onChange}
      placeholder={props.placeholder}
      onKeyDown={props.onKeyDown}
    />
  );
}

export default TextInput;
