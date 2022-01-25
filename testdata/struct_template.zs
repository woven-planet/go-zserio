package struct_template;

// Field is a lovely thing!
struct Field<T,Y>
{
    T value;
    Y other;
};

// Address, so I know where your house lives
struct Address
{
   Field<string,string> something;
};
