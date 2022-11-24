type Props = {
  text: string;
  onClick?:
  | ((event: React.MouseEvent<HTMLButtonElement, MouseEvent>) => void)
  | undefined;
};

function ButtonPrimary(props: Props) {
  return (
    <button
      className="bg-gradient-to-b from-[#3AB75B] to-[#066138] py-4 px-14 rounded-2xl"
      onClick={props.onClick}
    >
      <span className="text-white font-bold text-md">{props.text}</span>
    </button>
  );
}

export default ButtonPrimary;
