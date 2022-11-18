type Props = {
  placeholder: string;
  onChange: (event: React.ChangeEvent<HTMLInputElement>) => void;
};

function TextInput(props: Props) {
  return (
    <input
      className="py-4 px-3.5 rounded-2xl"
      onChange={props.onChange}
      placeholder={props.placeholder}
    />
  );
}

export default TextInput;
