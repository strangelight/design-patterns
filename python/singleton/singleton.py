#Singleton class
class Singleton:
    instance = None

    class __SingletonPriv:
        name = ""

        def instance(self):
            return repr(Singleton.instance)

    @staticmethod
    def GetInstance():
        if not Singleton.instance:
            Singleton.instance = Singleton.__SingletonPriv()
        return Singleton.instance

    def __init__(self):
        raise Exception('A singleton should not be instantiated! Use GetInstance method instead.')


########## usage:
x = Singleton.GetInstance()
x.name = "maicon"
print(x.name) #output: maicon

y = Singleton.GetInstance()
y.name = "patty"
print(y.name) #output: patty

print(x.name) #output: patty again

print(x.instance())
print(y.instance())
print(x.instance() == y.instance()) #output: true!

# RAISE AN EXCEPTION!!
z = Singleton()
