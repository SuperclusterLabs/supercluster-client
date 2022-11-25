type Props = {
  multiple?: boolean | undefined,
  onChange: (event: React.ChangeEvent<HTMLInputElement>) => void;
}

function Dropzone(props: Props) {
  return (
    <div className="flex items-center justify-center w-full">
      <label className="flex flex-col h-screen items-center justify-center w-full v-screen h-64 border-2 border-gray-300 border-dashed rounded-lg cursor-pointer bg-l-gray-50 hover:bg-l-gray-100">
        <div className="flex flex-col items-center justify-center pt-5 pb-6">
          <svg aria-hidden="true" className="w-10 h-10 mb-3 text-l-slateblue-700" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"></path></svg>
          <p className="mb-2 text-xl text-l-slateblue-700">Click to upload or drag and drop</p>
        </div>
        <input multiple={props.multiple} id="dropzone-file" type="file" className="hidden" onChange={props.onChange} />
      </label>
    </div>
  )
}

export default Dropzone;
