type Props = {
  text: string;
  onClick?:
  | ((event: React.MouseEvent<HTMLButtonElement, MouseEvent>) => void)
  | undefined;
};

function ButtonSecondary(props: Props) {
  return (
    <button
      className="ml-4 bg-gradient-to-b from-[#6A6D7C] to-[#3E3F4B] py-4 px-14 rounded-2xl"
      onClick={props.onClick}
    >
      <span className="text-white font-bold text-md">{props.text}</span>
    </button>
  );
}

export default ButtonSecondary;
