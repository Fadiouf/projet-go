
Construction de l'api :

1°) le dossier handlers

a°) handlers/auths.go

Ce code contient un ensemble de fonctions et de structures de gestion de l'authentification utilisateur dans une application Go, en utilisant le framework Echo et le package GORM pour interagir avec une base de données.

La structure AuthHandler est utilisée pour gérer les fonctions de création et de récupération des jetons d'authentification des utilisateurs.

La fonction Authenticate gère la création d'un nouveau jeton d'authentification en utilisant les informations fournies par l'utilisateur dans la requête HTTP.

La fonction AuthsHandler initialiser les routes HTTP nécessaire pour gérer l'authentification, dans ce cas, la route POST qui gère la création d'un nouveau jeton d'authentification.


b°) handlers/groups.go 

Ce code définit les gestionnaires  pour les itinéraires liés à la gestion des groupes.

Le package "handlers" une structure "GroupHandler" avec une référence à la base de données (DB) et des fonctions pour gérer les requêtes HTTP correspondant aux différentes routes de l'API.

La fonction "CreateGroup" gère la création d'un nouveau groupe, en liant les données de la requête à une structure de type "models.Group".

La fonction "GetGroups" retourne tous les groupes existants, récupérés depuis la base de données.

La fonction "UpdateGroup" met à jour un groupe existant, en récupérant le groupe correspondant à l'ID passé en paramètre, puis en mettant à jour ses champs avec les données de la requête.

La fonction "DeleteGroup" supprime un groupe existant, en récupérant le groupe correspondant à l'ID passé en paramètre, puis en le supprimant de la base de données.

Enfin, la fonction "GroupsHandler" initialise les routes correspondantes à ces fonctions en utilisant le framework "echo"


c°) handlers/roles.go

Ce code contient un gestionnaire de requêtes pour gérer les rôles dans une application Go. Il utilise le framework Echo pour créer des routes et gère les requêtes HTTP pour créer, lire, mettre à jour et supprimer des rôles dans la base de données.

Le code contient une structure RoleHandlerqui une instance gorm.DB pour accéder à la base de données.

Les fonctions CreateRole, GetRoles, UpdateRole, et DeleteRole sont des gestionnaires pour les opérations CRUD (Create, Read, Update, Delete) sur les rôles. Chacune de ces fonctions prend en entrée un objet echo.Context, qui représente la requête HTTP entrante, et renvoie une réponse HTTP. 
La fonction CreateRole crée un nouveau rôle en analysant les données JSON envoyées avec la requête HTTP. La fonction GetRoles renvoie tous les rôles stockés dans la base de données. 
La fonction UpdateRole met à jour un rôle existant en analysant les données JSON envoyées avec la requête HTTP et en appliquant les changements à l'objet rôle correspondant dans la base de données. Enfin, la fonction DeleteRole supprime un rôle existant de la base de données.

La fonction RolesHandler initialise les routes pour les gestionnaires de rôles en utilisant les méthodes HTTP GET, POST, PUT et DELETE.


d°) handlers/users.go

Ce code est un paquetage de gestionnaires pour le serveur HTTP écrit en Go. Les gestionnaires implémentent des fonctions pour les opérations CRUD (création, lecture, mise à jour et suppression) sur les utilisateurs.

Plus précisément, ce paquet contient les fonctions suivantes :

CreateUser : crée un nouvel utilisateur avec un mot de passe haché à partir du mot de passe brut fourni.
GetUsers : renvoie la liste de tous les utilisateurs, sans les mots de passe.
UpdateUser : met à jour les informations d'un utilisateur existant.
DeleteUser : supprime un utilisateur existant.
Il utilise le framework web Echo, qui fournit un moyen facile et rapide de créer des applications web en Go.


2°) le dossier models

a°) models/authtoken

Ce code est écrit en Go et concerne la mise en œuvre d'un middleware d'authentification basé sur JSON Web Token (JWT) pour une application Web.

Le package models contient deux structs de données : AuthTokenet AuthPayload. La première représente un token d'authentification, et la seconde représente la charge utile (payload) d'un token JWT.

Le middleware d'authentification ( authMiddleware) est appelé à chaque fois qu'une demande est envoyée à une route protégée, et il empêche si la demande contient un jeton JWT valide. Si le jeton est valide, il extrait le userID de la charge utile et le renvoie, sinon il renvoie une erreur.

La fonction createToken est utilisée pour créer un token JWT pour un utilisateur donné. Il définit la durée de validité du token et signe le token avec une clé secrète.

La fonction handleAuth est utilisée pour gérer le processus d'authentification. Elle valide les informations d'identification de l'utilisateur et, si elles sont valides, elle crée un token JWT en utilisant la fonction createToken et renvoie le token dans une réponse JSON.

Enfin, la fonction AuthsRoutes initialise les routes de l'application. Dans ce cas, il y a une seule route POST /auth qui appelle la fonction handleAuth


b°) models/group
Ce code est un package Go qui définit un modèle de groupe ( Group) et des fonctions de gestion de ces groupes à travers les routes HTTP ( handleGetGroups, handleCreateGroup, handleUpdateGroup, handleDeleteGroup, GroupsRoutes).

Le modèle Groupa plusieurs champs, notamment un ID(identifiant unique), un Name(nom), un Parent_Group_ID(identifiant du groupe parent), un tableau d'identifiants Child_Group_IDspour les groupes enfants, des champs de date de création ( Created_At), de dernière mise à jour ( Updated_At) et de suppression ( Deleted_At), et enfin des relations de plusieurs-à-plusieurs ( many2many) avec les modèles Useret Role.

Les fonctions handleGetGroups, handleCreateGroup, handleUpdateGroupet handleDeleteGroupsont des fonctions qui gèrent respectivement les requêtes HTTP GET, POST, PUT et DELETE pour les groupes. Elles prennent une requête HTTP ( echo.Context) en entrée, effectuent des opérations sur la base de données en fonction de la requête, puis retournent une réponse HTTP. Par exemple, handleGetGroupsrenvoyer tous les groupes existants en réponse HTTP JSON.

Enfin, la fonction GroupsRoutesinitialise les routes HTTP pour les groupes en utilisant la bibliothèque Echo. Elle crée un groupe de routes pour les groupes, définit les fonctions de gestion correspondantes pour chaque méthode HTTP et les associe aux routes correspondantes ( /groups, /groups/:id, etc.).



c°) models/roles
Ce code définit un modèle de rôle pour l'application web. Le modèle de rôle a un ID, un nom, une description, une date de création, une date de mise à jour, une date de suppression, une liste d'utilisateurs et une liste de groupes. Il y a également plusieurs fonctions pour manipuler les données de ce modèle.

La première fonction est handleGetRoles, qui récupère tous les rôles de la base de données et les renvoie sous forme JSON en réponse à une requête HTTP GET. La deuxième fonction handleCreateRole crée un nouveau rôle dans la base de données à partir du corps de la requête HTTP POST. 

La troisième fonction, handleUpdateRole récupère un rôle spécifique de la base de données à partir de son ID, met à jour ses champs avec les données du corps de la requête HTTP PUT, puis sauvegarde les changements dans la base de données. 

La dernière fonction, handleDeleteRole récupère un rôle spécifique de la base de données à partir de son ID et le supprime de la base de données.

Enfin, la fonction RolesRoutes initialise les routes pour les fonctions de manipulation de rôles en utilisant le framework Echo, en prévoyant un groupe de routes pour les rôles et en leur affectant les fonctions correspondantes pour les méthodes HTTP GET, POST, PUT et DELETE


d°) models/users
Ce code est une implémentation du serveur Web utilisant le framework "Echo" en Go. Le serveur implémente des routes pour gérer les utilisateurs stockés dans une base de données.

Le code commence par définir la structure de données pour l'utilisateur, qui est basé sur le modèle "gorm". Cette structure contient des informations sur l'utilisateur, telles que son nom, son email, son mot de passe, ses rôles, ses groupes et ses jetons d'authentification.

Ensuite, quatre fonctions sont définies pour gérer les requêtes HTTP pour les utilisateurs :

handleGetUsers récupérer tous les utilisateurs à partir de la base de données et renvoyer un code de réponse HTTP 200 (OK) avec les utilisateurs au format JSON.

handleCreateUser crée un nouvel utilisateur à partir du corps de la requête HTTP et l'ajoute à la base de données. Si la création a réussi, elle renvoie un code de réponse HTTP 200 (OK) avec les informations de l'utilisateur au format JSON.

handleUpdateUser Met à jour un utilisateur existant dans la base de données en fonction de l'ID de l'utilisateur fourni dans l'URL. Il récupère d'abord l'utilisateur à partir de la base de données, puis met à jour ses informations à partir du corps de la requête HTTP et enregistre les modifications dans la base de données. Si la mise à jour a réussi, elle renvoie un code de réponse HTTP 200 (OK) avec les informations de l'utilisateur au format JSON.

handleDeleteUser supprimer un utilisateur existant de la base de données en fonction de l'ID de l'utilisateur fourni dans l'URL. Il récupère d'abord l'utilisateur à partir de la base de données, puis le supprime de la base de données. Si la suppression a réussi, elle renvoie un code de réponse HTTP 204 (No Content).

Enfin, la fonction UsersRoutes est définie pour initialiser les routes pour la gestion des utilisateurs. Cette fonction prend un objet Echo en paramètre et définit quatre routes pour les quatre fonctions définies précédemment. Ces routes sont des sous-routes de /users et sont accessibles en tant que /users, /users/:id, etc.


e°) models/setup

Ce code contient une fonction SetupDB qui permet de se connecter à une base de données PostgreSQL en utilisant les informations de connexion intégrées dans des variables d'environnement.

La fonction utilise la bibliothèque GORM pour l'ORM, ainsi que le pilote PostgreSQL pour GORM ( gorm.io/driver/postgres).

Une fois la connexion établie, la fonction appelle AutoMigratesur plusieurs modèles, y compris User, Group, Role, AuthToken, et RefreshToken, pour s'assurer que les tables correspondantes existent dans la base de données. Cela permet de s'assurer que la base de données est synchronisée avec les définitions de modèle.

Le résultat de la fonction est un pointeur vers une instance gorm.DBqui peut être utilisée pour interagir avec la base de données. Le pointeur est stocké dans une variable globale DB pour être utilisé dans d'autres parties de l'application.


3°) main.go

Ce code est la fonction main() du programme. Il est responsable de la configuration de l'application et de l'exécution du serveur.

Tout d'abord, la fonction charge les variables d'environnement depuis un fichier .env à l'aide de la bibliothèque GoDotEnv. Les variables d'environnement sont utilisées pour configurer la connexion à la base de données et le port du serveur.

Ensuite, la fonction initialise une connexion à la base de données en appelant la fonction SetupDB() du package models. Si la connexion n'est pas établie avec succès, la fonction panique.

Après cela, la fonction crée une instance de l'application Echo, qui est un framework web léger pour Go. Ensuite, elle ajoute deux middlewares : le middleware Logger() qui enregistre les requêtes HTTP et les réponses HTTP, et le middleware Recover() qui capture les erreurs pour empêcher le serveur de s'arrêter.

Ensuite, la fonction initialise les routes en appelant les fonctions *Routes() du package models. Ces fonctions génèrent les points de terminaison de l'API et les fonctions qui gèrent ces points de terminaison.

Après cela, la fonction initialise les gestionnaires (handlers) en appelant les fonctions *Handler() du package handlers. Ces fonctions génèrent les fonctions qui gèrent les requêtes HTTP.

Enfin, la fonction démarrer le serveur sur le port spécifié par les variables d'environnement ou sur le port 8080 si aucune variable d'environnement n'est supprimée. Si le serveur ne peut pas démarrer, la fonction panique.



Exécution de l'API :

Prérequis : 
. Avoir une base de donnée postgres d'installée et configurée.

Pour tester le code source de l'api, il faut télécharger le dossier /projet-go 
puis l'ouvrir avec Visual Studio Code ou autre éditeur de code.

Ensuite aller sur le fichier .env pour modifier les informations de connexion en fonction des informations de votre base de donnée.


Enfin pour l'éxécution du code, il faut se mettre sur la racine du dossier projet-go puis lancer la commande go run main.go.

